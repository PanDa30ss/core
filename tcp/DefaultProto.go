package tcp

import "github.com/PanDa30ss/core/message"

type PingReq struct {
	message.Package
}

func MakePingReq() *PingReq {
	ret := &PingReq{}
	ret.Dump = func(msg *message.Message) {
	}
	ret.Load = func(msg *message.Message) {
	}
	return ret
}

type PingAck struct {
	message.Package
}

func MakePingAck() *PingAck {
	ret := &PingAck{}
	ret.Dump = func(msg *message.Message) {
	}
	ret.Load = func(msg *message.Message) {
	}
	return ret
}
