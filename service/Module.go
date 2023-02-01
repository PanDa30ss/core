package service

import log "github.com/PanDa30ss/core/logManager"

type IModule interface {
	Init()
	Start() bool
	Stop()
	CheckStart() bool
	IsStarted() bool
	SetStart(started bool)
	GetName() string
	SetName(name string)
	Initial()
}

type Module struct {
	started bool
	name    string
}

func (this *Module) Init() {

}

func (this *Module) Start() bool {
	log.Info("Module Start", this.GetName())
	return true
}

func (this *Module) Initial() {
	log.Info("Module Initial", this.GetName())
	return
}

func (this *Module) CheckStart() bool {
	return true
}

func (this *Module) Stop() {

}

func (this *Module) IsStarted() bool {
	return this.started
}

func (this *Module) SetStart(started bool) {
	this.started = started

}
func (this *Module) GetName() string {
	return this.name
}

func (this *Module) SetName(name string) {
	this.name = name
}
