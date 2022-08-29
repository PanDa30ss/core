package tcp

type ISessionS interface {
	ISession
	// setConv(conv int)
	// GetConv() int
	GetBank() ISessionSBank
	bindBank(bank ISessionSBank)
}

type SessionS struct {
	Session
	bank ISessionSBank
	// conv int
}

func (this *SessionS) GetBank() ISessionSBank {
	return this.bank
}

func (this *SessionS) bindBank(bank ISessionSBank) {
	this.bank = bank
}

// func (this *SessionS) GetConv() int {
// 	return this.conv
// }

// func (this *SessionS) setConv(conv int) {
// 	this.conv = conv
// }

func (this *SessionS) close() {
	this.Session.close()
	// if this.bank != nil {
	// 	this.bank.removeSession(this.conv)
	// }
}

func (this *SessionS) open() bool {
	return this.Session.open()
}
