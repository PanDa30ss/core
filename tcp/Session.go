package tcp

import (
	"encoding/binary"
	"io"
	"net"

	log "github.com/PanDa30ss/core/logManager"
	"github.com/PanDa30ss/core/message"
	"github.com/PanDa30ss/core/service"
)

type ISession interface {
	run() bool
	SendMessage(msg *message.Message)
	CallParseMessage(data []byte)
	open() bool
	IsOpen() bool
	close()
	bind(conn net.Conn)
	getError() int
	setPingData(timeData *[]int)
	BindPingManager(timeData *[]int)
	SetInstance(session ISession)
	// PostClose()
	OnOpen()
	OnClose()
}

type Session struct {
	IP          string
	Port        uint32
	isOpen      bool
	conn        net.Conn
	sendBuffer  *sendBuffer
	errorCode   int
	factory     *commandFactory
	pingManager *pingManager
	instance    ISession
}

func (this *Session) SetInstance(session ISession) {
	this.instance = session
}

func (this *Session) bind(conn net.Conn) {
	this.conn = conn
}

func (this *Session) BindPingManager(timeData *[]int) {
	this.pingManager = makePingManager(timeData, this)
}

func (this *Session) open() bool {
	if this.isOpen {
		return true
	}

	this.isOpen = true
	this.sendBuffer = newSendBuffer(this)
	this.errorCode = 0
	this.pingManager.open()
	go this.startRead()

	return true
}

func (this *Session) IsOpen() bool {
	return this.isOpen
}

func (this *Session) close() {
	if !this.isOpen {
		return
	}
	log.Info("session close %v", this.errorCode)
	this.isOpen = false
	if this.conn != nil {
		this.conn.Close()
	}

}

func (this *Session) setPingData(timeData *[]int) {
	this.pingManager.timeData = timeData
}

func (this *Session) setError(code int) {
	this.errorCode = code
}

func (this *Session) getError() int {
	return this.errorCode
}

func (this *Session) startRead() {
	defer this.setError(ERROR_RECVFAILED)
	for {
		buf := make([]byte, 2)
		// _, err := this.conn.Read(buf)
		_, err := io.ReadFull(this.conn, buf)
		if err != nil {
			this.setError(ERROR_RECVFAILED)
			return
		}
		msgLen := binary.BigEndian.Uint16(buf)
		dataBuf := make([]byte, msgLen)
		_, errData := io.ReadFull(this.conn, dataBuf[2:])
		if errData != nil {
			this.setError(ERROR_RECVFAILED)
			return
		}
		copy(dataBuf[0:], buf)
		cmd := makeSessionMessageCommand(this.instance, dataBuf)
		service.Post(cmd)
	}
}

func (this *Session) startSend() {
	buf, size := this.sendBuffer.takeSendData()
	if buf == nil {
		return
	}
	go this.handleSend(buf, size)
}

func (this *Session) handleSend(buf []byte, size int) {
	// for size > 0 {
	_, err := this.conn.Write(buf[:size])
	if err != nil {
		this.setError(ERROR_SENDFAILED)
		return
	}
	// if n > size {
	// 	this.setError(ERROR_SENDFAILED)
	// 	return
	// }
	// size -= n
	// }
	this.sendBuffer.popData()
	this.startSend()
}

func (this *Session) run() bool {
	if !this.isOpen {
		return false
	}
	if !this.pingManager.run() {
		this.setError(ERROR_PINGTIMEOUT)
	}
	return this.errorCode == 0
}

func (this *Session) SetFactory(fid int) {
	this.factory = getCommandStorage().factorys[fid]
}

func (this *Session) send(msg *message.Message) {
	buff := msg.GetData()
	if this.sendBuffer == nil {
		return
	}
	if !this.sendBuffer.pushData(buff) {
		this.setError(ERROR_SENDBUFFER)
	} else {
		this.startSend()
	}
}

func (this *Session) SendMessage(msg *message.Message) {
	this.send(msg)
}

func (this *Session) ParseMessage(msg *message.Message) (result bool) {

	this.pingManager.onPing()
	return this.factory.funcs[msg.ID()](this.instance, msg)
}

func (this *Session) CallParseMessage(data []byte) {

	msg := message.MakeMessage()
	msg.SetData(data)

	if !this.ParseMessage(msg) {
		this.setError(ERROR_PARSEMESSAGE)
	}

}

func (this *Session) OnOpen() {
}

func (this *Session) OnClose() {
}
