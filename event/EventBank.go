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
