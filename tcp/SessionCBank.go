package tcp

import (
	"container/list"
	"core/service"
	"sync"
	"time"
)

// "service"

type ISessionCBank interface {
	Init(bank ISessionCBank)
	Close()
	Start()
	AddConnectSession(session ISessionC)
	Running() chan struct{}
}

type SessionCBank struct {
	mutex               sync.Mutex
	running             chan struct{}
	connectorSessions   *list.List
	doConnectorSessions *list.List
	sessionList         *list.List
}

func (this *SessionCBank) Init() {
	this.mutex = sync.Mutex{}
	// this.sessionList = list.New()
	// this.connectorSessions = list.New()
	// this.doConnectorSessions = list.New()
}

func (this *SessionCBank) Start() {
	this.running = make(chan struct{})
	// go this.goRun()
}

// func (this *SessionCBank) getConnectSessions() {
// 	this.mutex.Lock()
// 	defer this.mutex.Unlock()
// 	this.connectorSessions, this.doConnectorSessions = this.doConnectorSessions, this.connectorSessions

// }

// func (this *SessionCBank) AddConnectSession(session ISessionC) {
// 	this.mutex.Lock()
// 	defer this.mutex.Unlock()
// 	this.doConnectorSessions.PushBack(session)
// }

func (this *SessionCBank) Running() chan struct{} {
	return this.running
}

func (this *SessionCBank) Close() {
	close(this.running)
}

// func (this *SessionCBank) goRun() {
// 	for {
// 		select {
// 		case <-this.Running():
// 			log.Info("SessionCBank close")
// 			return
// 		default:
// 			this.run()
// 			time.Sleep(time.Millisecond * 1)
// 		}
// 	}
// }

// func (this *SessionCBank) run() {
// 	this.getConnectSessions()
// 	for e := this.doConnectorSessions.Front(); e != nil; e = e.Next() {
// 		this.sessionList.PushBack(this.doConnectorSessions.Remove(e).(ISessionC))
// 	}

// 	for e := this.sessionList.Front(); e != nil; e = e.Next() {
// 		session := e.Value.(ISessionC)
// 		switch session.GetStatus() {
// 		case WaitConnect:
// 			session.Connect()
// 		case Fail:
// 			{
// 				session.Reset()
// 				session.Connect()
// 			}
// 		case Connected:
// 			{
// 				if session.open() {
// 					service.Post(makeSessionOpenCommand(session))
// 				} else {
// 					session.close()
// 				}
// 			}
// 		case Open:
// 			if !session.run() {
// 				session.close()
// 				service.Post(makeSessionCloseCommand(session))
// 			}
// 		}

// 	}
// }

func (this *SessionCBank) AddConnectSession(session ISessionC) {
	session.SetInstance(session)
	go this.goSessionRun(session)
}

func (this *SessionCBank) goSessionRun(session ISessionC) {
	for !session.IsRemoved() {
		switch session.GetStatus() {
		case WaitConnect:
			session.Connect()
		case Fail:
			{
				session.Reset()
				session.Connect()
			}
		case Connected:
			{
				if session.open() {
					service.Post(makeSessionOpenCommand(session))
				} else {
					session.close()
				}
			}
		case Open:
			if !session.run() {
				session.close()
				service.Post(makeSessionCloseCommand(session))
			}
		}
		time.Sleep(time.Millisecond * 500)
	}

	session.close()
	service.Post(makeSessionCloseCommand(session))
}
