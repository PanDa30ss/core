package tcp

import (
	"time"

	"github.com/PanDa30ss/core/message"
)

type pingManager struct {
	timeData  *[]int
	pingTime  time.Time
	pingCount int
	session   ISession
}

func (this *pingManager) onPing() {
	this.pingCount = 0
	this.setTime()
}
func (this *pingManager) ping() {
	req := MakePingReq()
	m := message.MakeMessageWithPackage(PING_REQ, req)
	this.session.SendMessage(m)
}
func (this *pingManager) run() bool {
	if this.pingCount >= len(*this.timeData) {
		return false
	}
	if time.Now().After(this.pingTime) {
		this.pingCount++
		if this.pingCount >= len(*this.timeData) {
			return false
		}
		this.ping()
		this.setTime()
	}
	return true
}
func (this *pingManager) open() {
	this.setTime()
}

func (this *pingManager) Reset() {
	this.pingCount = 0
	this.setTime()
}

func (this *pingManager) setTime() {
	this.pingTime = time.Now().Add(time.Duration((*this.timeData)[this.pingCount]) * time.Second)
}

func makePingManager(timeData *[]int, session ISession) *pingManager {
	var ret = &pingManager{}
	ret.timeData = timeData
	ret.session = session
	ret.Reset()
	return ret
}
