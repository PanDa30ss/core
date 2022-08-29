package service

func Init() {
	getInstance().init()
}

func Start() bool {
	return getInstance().start()
}

func Stop() {
	getInstance().stop()
}

func Post(cmd ICommand) {
	go func() {
		getInstance().cmdChan <- cmd
	}()
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
