package event

import (
	"container/list"
	"sync"
)

type eventFunc func(params ...interface{})

type eventBank struct {
	// service.Module
	curUID     int
	handlerMap map[int]*list.List
	// nameMap    map[string]int
}

var once sync.Once
var instance *eventBank

func getInstance() *eventBank {
	once.Do(func() {
		instance = makeEventBank()
	})
	return instance
}

func makeEventBank() *eventBank {
	eb := &eventBank{}
	eb.curUID = 0
	eb.handlerMap = make(map[int]*list.List)
	// eb.nameMap = make(map[int]string)
	return eb
}

func GetEventID() int {
	this := getInstance()
	this.curUID++
	this.handlerMap[this.curUID] = list.New()
	// this.nameMap[name] = this.curUID
	return this.curUID
}

func MakeEvent(uid int, params ...interface{}) *event {
	e := &event{uid, params}
	return e
}

func RegisterEventHandler(uid int, f eventFunc) bool {
	l, ok := getInstance().handlerMap[uid]
	if !ok {
		panic("event uid error")
	}
	l.PushBack(f)
	return true
}

func DispatchEvent(ev *event) {
	l, ok := getInstance().handlerMap[ev.uid]
	if !ok {
		return
	}
	for e := l.Front(); e != nil; e = e.Next() {
		e.Value.(eventFunc)(ev.params...)
	}
	return
}
