package event

import "container/list"

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
