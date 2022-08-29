package tcp

func GetSessionUID() int {
	return getCommandStorage().getSessionStorageUID()
}

func RegisterCMD(uid int, cmd uint16, f cmdFunc) bool {
	return getCommandStorage().registerCMD(uid, cmd, f)
}
