package service

import (
	"fmt"

	log "github.com/PanDa30ss/core/logManager"
)

type IModule interface {
	Init()
	Start() bool
	Stop()
	CheckStart(modules map[string]IModule) bool
	IsStarted() bool
	SetStart(started bool)
	GetName() string
	SetName(name string)
	Initial()
	CheckInitial(modules map[string]IModule) bool
	IsInitialed() bool
	SetInitial(initialed bool)
}

type Module struct {
	initialed bool
	started   bool
	name      string
	depends   []string
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

func (this *Module) CheckStart(modules map[string]IModule) bool {
	for i := 0; i < len(this.depends); i++ {
		if m, ok := modules[this.depends[i]]; ok {
			if !m.IsStarted() {
				return false
			}
		} else {
			e := fmt.Sprintf("error module %v", this.depends[i])
			panic(e)
		}

	}
	return true
}

func (this *Module) CheckInitial(modules map[string]IModule) bool {
	for i := 0; i < len(this.depends); i++ {
		if m, ok := modules[this.depends[i]]; ok {
			if !m.IsInitialed() {
				return false
			}
		} else {
			e := fmt.Sprintf("error module %v", this.depends[i])
			panic(e)
		}

	}
	return true
}

func (this *Module) Stop() {

}

func (this *Module) IsInitialed() bool {
	return this.initialed
}

func (this *Module) SetInitial(initialed bool) {
	this.initialed = initialed
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

func (this *Module) SetDepends(depends []string) {
	this.depends = depends
}
