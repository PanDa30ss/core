package tcp

import (
	log "core/logManager"
	"net"
)

type ISessionC interface {
	ISession
	SetAddress(addr string)
	Connect()
	Reset()
	GetStatus() int
	IsRemoved() bool
}

type SessionC struct {
	Session
	connectStatus int
	addr          string
	removed       bool
	// bank          ISessionCBank
	// e             *list.Element
}

func (this *SessionC) GetStatus() int {
	return this.connectStatus
}

func (this *SessionC) SetAddress(addr string) {
	this.addr = addr
}

func (this *SessionC) Connect() {
	if this.connectStatus != WaitConnect {
		return
	}
	this.connectStatus = Connecting

	go this.tryConnect()
}

func (this *SessionC) tryConnect() {
	if this.addr == "" {
		return
	}
	conn, err := net.Dial("tcp", this.addr)
	if err != nil {
		log.Info(err)
		this.connectStatus = Fail
		return
	}
	this.connectStatus = Connected
	this.conn = conn

}

func (this *SessionC) Reset() {
	this.pingManager.Reset()
	this.setError(0)
	this.connectStatus = WaitConnect
}

func (this *SessionC) close() {
	this.Session.close()
	this.Reset()
}

func (this *SessionC) open() bool {
	this.connectStatus = Open
	return this.Session.open()
}

func (this *SessionC) Remove() {

	this.removed = true
}
func (this *SessionC) IsRemoved() bool {
	return this.removed
}
