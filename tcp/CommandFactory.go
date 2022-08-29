package tcp

import (
	"core/message"
)

type cmdFunc func(s ISession, msg *message.Message) bool

type commandFactory struct {
	funcs [CMD_SIZE]cmdFunc
}

func (this *commandFactory) addFunc(cmd uint16, f cmdFunc) {
	this.funcs[cmd] = f

}

func defaultCmdFunc(s ISession, msg *message.Message) bool {
	return false
}

func newCommandFactory() *commandFactory {
	ft := &commandFactory{}
	for i := 0; i < CMD_SIZE; i++ {
		ft.funcs[i] = defaultCmdFunc
	}
	ft.funcs[PING_ACK] = onPingAck
	ft.funcs[PING_REQ] = onPingReq
	return ft
}

func onPingReq(s ISession, msg *message.Message) bool {
	ack := MakePingAck()
	m := message.MakeMessageWithPackage(PING_ACK, ack)
	s.SendMessage(m)
	return true
}

func onPingAck(s ISession, msg *message.Message) bool {
	return true
}
