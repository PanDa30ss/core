package message

import (
	"bytes"
	"encoding/binary"

	"github.com/golang/protobuf/proto"
)

type Message struct {
	data []byte
	buf  *bytes.Buffer
}

func (this *Message) ID() uint16 {
	this.ReadUInt16()
	id := this.ReadUInt16()
	return id
}

func (this *Message) SetID(id uint16) {
	this.Write(uint16(0))
	this.Write(id)
}

// func (this *Message) ResetID(id uint16) {
// 	this.msgID = id
// 	buf := bytes.NewBuffer([]byte{})
// 	binary.Write(buf, binary.BigEndian, &this.msgID)
// 	copy(this.buf.Bytes()[2:], buf.Bytes())
// 	binary.BigEndian.Uint32()
// }

func (this *Message) SetData(data []byte) {
	this.data = data
	this.buf = bytes.NewBuffer(this.data)
}

func (this *Message) GetData() []byte {
	return this.data
}

func (this *Message) GetBody() []byte {
	return this.data[4:]
}

func (this *Message) GetSize() int {
	return this.buf.Len()
}

func (this *Message) Done() {
	this.WriteLength()
	this.data = this.buf.Bytes()
}

func (this *Message) WriteLength() {
	binary.BigEndian.PutUint16(this.buf.Bytes(), uint16(this.buf.Len()))
}

func (this *Message) Marshal(p interface{}) error {
	data, err := proto.Marshal(p.(proto.Message))
	if err != nil {
		return err
	}
	this.WriteBytes(data)
	return err
}

func (this *Message) Unmarshal(p interface{}) error {
	data := this.ReadBytes()
	err := proto.UnmarshalMerge(data, p.(proto.Message))
	return err
}

func (this *Message) Write(data interface{}) {
	binary.Write(this.buf, binary.BigEndian, data)
}
func (this *Message) Read(data interface{}) {
	binary.Read(this.buf, binary.BigEndian, data)
}

func (this *Message) ReadUInt64() uint64 {
	var data uint64
	binary.Read(this.buf, binary.BigEndian, &data)
	return data
}

func (this *Message) ReadUInt32() uint32 {
	var data uint32
	binary.Read(this.buf, binary.BigEndian, &data)
	return data
}

func (this *Message) ReadUInt16() uint16 {
	var data uint16
	binary.Read(this.buf, binary.BigEndian, &data)
	return data
}
func (this *Message) ReadInt64() int64 {
	var data int64
	binary.Read(this.buf, binary.BigEndian, &data)
	return data
}

func (this *Message) ReadInt32() int32 {
	var data int32
	binary.Read(this.buf, binary.BigEndian, &data)
	return data
}

func (this *Message) ReadInt16() int16 {
	var data int16
	binary.Read(this.buf, binary.BigEndian, &data)
	return data
}

func (this *Message) WriteString(s string) {
	var b byte = byte(0)
	this.buf.WriteString(s)
	this.buf.WriteByte(b)
}

func (this *Message) ReadString(s *string) {
	var b byte = byte(0)
	*s, _ = this.buf.ReadString(b)
	*s = (*s)[:len(*s)-1]
}
func (this *Message) WriteUInt8Array(a []uint8) {
	var l uint16 = uint16(len(a))
	binary.Write(this.buf, binary.BigEndian, &l)
	binary.Write(this.buf, binary.BigEndian, &a)
}

func (this *Message) ReadUInt8Array(a []uint8) {
	var l uint16
	this.Read(&l)
	a = make([]uint8, l)
	binary.Read(this.buf, binary.BigEndian, a)
}

func (this *Message) WriteUInt32Array(a []uint32) {
	var l uint16 = uint16(len(a))
	binary.Write(this.buf, binary.BigEndian, &l)
	binary.Write(this.buf, binary.BigEndian, &a)
}

func (this *Message) ReadUInt32Array(a []uint32) {
	var l uint16
	this.Read(&l)
	a = make([]uint32, l)
	binary.Read(this.buf, binary.BigEndian, a)
}

func (this *Message) WriteArray(a []interface{}) {
	var l uint16 = uint16(len(a))
	binary.Write(this.buf, binary.BigEndian, &l)
	for i := uint16(0); i < l; i++ {
		a[i].(IPackage).GetDumpFunc()(this)
	}
}

func (this *Message) ReadArray(a []IPackage) {
	var l uint16
	this.Read(&l)
	a = make([]IPackage, l)

	for i := uint16(0); i < l; i++ {
		a[i].GetLoadFunc()(this)
	}

}

func (this *Message) WriteBytes(s []byte) {
	var l uint16 = uint16(len(s))
	binary.Write(this.buf, binary.BigEndian, &l)
	this.buf.Write(s)
}

func (this *Message) ReadBytes() []byte {
	var l uint16
	this.Read(&l)
	s := make([]byte, l)
	this.buf.Read(s)
	return s
}
func (this *Message) WritePackage(a IPackage) {
	a.GetDumpFunc()(this)
}

func (this *Message) ReadPackage(a IPackage) {
	a.GetLoadFunc()(this)
}
