package service

import (
	"sync"
	"time"
)

// var serviceQueueLen uint32 = 0x10000

type service struct {
	modules map[string]IModule
	// queue   *Queue
	cmdChan chan ICommand
}

var instance *service
var once sync.Once

func getInstance() *service {
	once.Do(func() {
		instance = newService()
	})
	return instance
}

func newService() *service {
	var this = service{}
	// this.queue = makeQueue()
	this.cmdChan = make(chan ICommand, 1)
	this.modules = make(map[string]IModule)
	// this.modules = make([]IModule, 0, 0)
	return &this
}

func (this *service) goRun() {
	for {
		select {
		case cmd := <-this.cmdChan:
			cmd.(ICommand).Execute()
		}

	}
	// for {
	// 	cmd := this.queue.deQueue()
	// 	cmd.Execute()
	// }
}

func (this *service) init() {
	for _, m := range this.modules {
		m.Init()
		m.SetStart(false)
	}
}

func (this *service) start() bool {
	timer := time.NewTimer(3 * time.Second)
	num := 0
	for num < len(this.modules) {
		select {
		case <-timer.C:
			return false
		default:
			{
				for _, m := range this.modules {
					if m.IsStarted() {
						continue
					}
					if !m.CheckStart() {
						continue
					}
					if !m.Start() {
						return false
					}
					m.SetStart(true)
					num++
				}
			}
		}
	}

	go this.goRun()
	return true
}

func (this *service) stop() {

	for _, m := range this.modules {
		m.Stop()
	}
}

func (this *service) selectModules(modules map[string]bool) bool {
	for name, _ := range modules {
		if _, ok := this.modules[name]; !ok {
			return false
		}
	}

	for name, _ := range this.modules {
		if _, ok := modules[name]; ok {
			continue
		}
		delete(this.modules, name)
	}
	return true
}

func Init() {
	getInstance().init()
}

func Start() bool {
	return getInstance().start()
}

func Stop() {
	getInstance().stop()
}

func GoPost(cmd ICommand) {
	go func() {
		getInstance().cmdChan <- cmd
	}()
	// getInstance().queue.enQueue(cmd)
}

func Post(cmd ICommand) {
	getInstance().cmdChan <- cmd
	// getInstance().queue.enQueue(cmd)
}

func RegisterModule(name string, m IModule) bool {
	m.SetName(name)
	if _, ok := getInstance().modules[m.GetName()]; ok {
		panic("duplicate module key")
	}
	getInstance().modules[m.GetName()] = m
	return true
}

func SelectModule(modules map[string]bool) bool {
	return getInstance().selectModules(modules)
}
