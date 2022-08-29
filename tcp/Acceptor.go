package tcp

import (
	"net"
	"time"

	log "github.com/PanDa30ss/core/logManager"
	"github.com/PanDa30ss/core/service"
)

type acceptor struct {
	bank ISessionSBank
	// listener *net.Listener
}

func (this *acceptor) start() {
	listener, err := net.Listen("tcp", this.bank.getAddr())
	if err != nil {
		log.Error("acceptor Listen error", err)
		panic(err)
	}
	for {
		select {
		case <-this.bank.Running():
			log.Info("acceptor close")
			return
		default:
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			var session = this.bank.makeSession(conn)

			if !session.open() {
				session.close()
				return
			}
			service.Post(makeSessionOpenCommand(session))
			go this.goSessionRun(session)
		}

	}
}

func (this *acceptor) goSessionRun(session ISessionS) {
	for session.run() {
		time.Sleep(time.Millisecond * 500)
	}
	session.close()
	service.Post(makeSessionCloseCommand(session))
}
