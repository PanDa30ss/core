package tcp

import "sync"

var sessionStorageUID int = 0

var storage *commandStorage
var onceStorage sync.Once

func getCommandStorage() *commandStorage {
	onceStorage.Do(func() {
		storage = &commandStorage{make(map[int]*commandFactory)}
	})
	return storage
}

type commandStorage struct {
	factorys map[int]*commandFactory
}

func (this *commandStorage) getSessionStorageUID() int {
	sessionStorageUID++
	this.factorys[sessionStorageUID] = newCommandFactory()
	return sessionStorageUID
}

func (this *commandStorage) registerCMD(uid int, cmd uint16, f cmdFunc) bool {
	factory, ok := this.factorys[uid]
	if !ok {
		return false
	}
	factory.addFunc(cmd, f)
	return true
}

func GetSessionUID() int {
	return getCommandStorage().getSessionStorageUID()
}

func RegisterCMD(uid int, cmd uint16, f cmdFunc) bool {
	return getCommandStorage().registerCMD(uid, cmd, f)
}
