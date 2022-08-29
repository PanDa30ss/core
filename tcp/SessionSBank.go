package tcp

import (
	"container/list" // "service"
	"net"
	"sync"
)

type ISessionSBank interface {
	BindAddr(addr string)
	getAddr() string
	Init(bank ISessionSBank)
	Close()
	Start()
	Running() chan struct{}
	SetMaxSession(max int)
	CreateSession() ISessionS
	// FindSession(conv int) ISessionS
	// removeSession(conv int)
	// addSession(session ISessionS)
	// getNextConv() int
	SetPingData(timeData *[]int)
	// addOpenSession(session ISessionS)
	makeSession(conn net.Conn) ISessionS
}

type sessionItem struct {
	session ISessionS
	it      *list.Element
}

type SessionSBank struct {
	mutex sync.Mutex
	// openSessions   *list.List
	// doOpenSessions *list.List
	running  chan struct{}
	acceptor *acceptor
	addr     string
	// sessionList *list.List
	// sessionMap  map[int]*sessionItem
	maxSessions int
	nextConv    int
	pingData    *[]int
}

func (this *SessionSBank) BindAddr(addr string) {
	this.addr = addr
}

func (this *SessionSBank) getAddr() string {
	return this.addr
}

func (this *SessionSBank) SetPingData(timeData *[]int) {
	this.pingData = timeData
}

func (this *SessionSBank) SetMaxSession(max int) {
	this.maxSessions = max
}

func (this *SessionSBank) Init(bank ISessionSBank) {
	this.mutex = sync.Mutex{}
	this.pingData = DefaultPingTimeData
	// this.openSessions = list.New()
	// this.doOpenSessions = list.New()
	// this.sessionList = list.New()
	// this.sessionMap = make(map[int]*sessionItem)
	this.acceptor = &acceptor{}
	this.acceptor.bank = bank
	this.nextConv = 0
	this.maxSessions = 1024
}

// func (this *SessionBank) BindAcceptor(bank ISessionBank) {
// }
// func (this *SessionSBank) getOpenSessions() {
// 	this.mutex.Lock()
// 	defer this.mutex.Unlock()
// 	this.openSessions, this.doOpenSessions = this.doOpenSessions, this.openSessions

// }

// func (this *SessionSBank) addOpenSession(session ISessionS) {
// 	this.mutex.Lock()
// 	defer this.mutex.Unlock()
// 	this.openSessions.PushBack(session)
// }

// func (this *SessionSBank) goRun() {
// 	for {
// 		select {
// 		case <-this.Running():
// 			log.Info("SessionSBank close")
// 			return
// 		default:
// 			this.run()
// 			time.Sleep(time.Millisecond * 1)
// 		}
// 	}
// }

// func (this *SessionSBank) run() {
// 	this.getOpenSessions()
// 	for e := this.doOpenSessions.Front(); e != nil; e = e.Next() {
// 		this.addSession(this.doOpenSessions.Remove(e).(ISessionS))
// 	}

// 	for e := this.sessionList.Front(); e != nil; e = e.Next() {
// 		session := e.Value.(*sessionItem).session
// 		if !session.run() {
// 			session.close()
// 			service.Post(makeSessionCloseCommand(session))
// 		}
// 	}

// }

func (this *SessionSBank) Running() chan struct{} {
	return this.running
}

func (this *SessionSBank) Close() {
	close(this.running)
}

func (this *SessionSBank) Start() {
	this.running = make(chan struct{})
	go this.acceptor.start()
	// go this.goRun()

}

func (this *SessionSBank) CreateSession() ISessionS {
	return nil
}

// func (this *SessionSBank) getNextConv() int {
// 	if this.sessionList.Len() >= this.maxSessions {
// 		return INVALID_CONV
// 	}

// 	for {
// 		if this.nextConv == INVALID_CONV {
// 			this.nextConv++
// 		}
// 		if _, ok := this.sessionMap[this.nextConv]; !ok {
// 			ret := this.nextConv
// 			this.nextConv++
// 			return ret
// 		}
// 		this.nextConv++
// 	}
// }

// func (this *SessionSBank) addSession(session ISessionS) {

// 	conv := this.getNextConv()
// 	if conv == INVALID_CONV {
// 		return
// 	}
// 	session.setConv(conv)

// 	if !session.open() {
// 		session.close()
// 		return
// 	}

// 	service.Post(makeSessionOpenCommand(session))

// 	item := &sessionItem{}
// 	item.session = session
// 	item.it = this.sessionList.PushBack(item)
// 	this.sessionMap[session.GetConv()] = item
// }

// func (this *SessionSBank) FindSession(conv int) ISessionS {
// 	if item, ok := this.sessionMap[conv]; ok {
// 		return item.session
// 	} else {
// 		return nil
// 	}
// }

// func (this *SessionSBank) removeSession(conv int) {
// 	if item, ok := this.sessionMap[conv]; ok {
// 		this.sessionList.Remove(item.it)
// 		delete(this.sessionMap, conv)
// 	}

// }

func (this *SessionSBank) makeSession(conn net.Conn) ISessionS {
	var session = this.acceptor.bank.CreateSession()
	session.bind(conn)
	session.bindBank(this.acceptor.bank)
	session.BindPingManager(this.pingData)
	session.SetInstance(session)
	return session

}
