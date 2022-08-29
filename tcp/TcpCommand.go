package tcp

import (
	"core/service"
)

type tcpOpenCommand struct {
	session ISession
}

func (this *tcpOpenCommand) Execute() {
	this.session.OnOpen()

}
func makeSessionOpenCommand(s ISession) service.ICommand {
	return &tcpOpenCommand{session: s}
}

type tcpCloseCommand struct {
	session ISession
}

func (this *tcpCloseCommand) Execute() {
	this.session.OnClose()
}

func makeSessionCloseCommand(s ISession) service.ICommand {
	return &tcpCloseCommand{session: s}
}

type tcpMessageCommand struct {
	session ISession
	data    []byte
}

func (this *tcpMessageCommand) Execute() {
	this.session.CallParseMessage(this.data)

}

func makeSessionMessageCommand(s ISession, d []byte) service.ICommand {
	return &tcpMessageCommand{session: s, data: d}
}
