package message

import (
	"bytes"
)

func MakeMessage() *Message {
	m := &Message{}
	m.data = make([]byte, 0)
	m.buf = bytes.NewBuffer([]byte{})
	return m
}

func MakeMessageWithPackage(id uint16, pkg IPackage) *Message {
	msg := MakeMessage()
	msg.SetID(id)
	pkg.GetDumpFunc()(msg)
	msg.Done()
	return msg
}

func MakePBMessage(id uint16, p interface{}) *Message {
	msg := MakeMessage()
	msg.SetID(id)
	msg.Marshal(p)
	msg.Done()
	return msg
}
