package tcp

import (
	"container/list"
	"sync"
)

type block struct {
	buff []byte
	size int
}
type sendBuffer struct {
	mutex        sync.Mutex
	session      *Session
	dataSize     int
	blocks       *list.List
	sendingBlock *block
}

func newSendBuffer(session *Session) *sendBuffer {
	s := &sendBuffer{}
	s.mutex = sync.Mutex{}
	s.session = session
	s.dataSize = 0
	s.blocks = list.New()
	s.sendingBlock = nil
	return s
}

func (this *sendBuffer) pushData(data []byte) bool {
	size := len(data)
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if this.dataSize+size > SEND_MAX_SIZE {
		return false
	}
	leftSize := size
	if this.blocks.Len() > 0 {
		b := this.blocks.Back().Value.(*block)
		space := BLOCK_SIZE - b.size
		if space >= 0 {
			var writeSize int
			if space > size {
				writeSize = size
			} else {
				writeSize = space
			}
			copy(b.buff[b.size:], data[:writeSize])
			b.size += writeSize
			leftSize -= writeSize
		}
	}
	if leftSize > 0 {
		b := &block{buff: make([]byte, BLOCK_SIZE), size: 0}
		this.blocks.PushBack(b)
		b.size += leftSize
		copy(b.buff, data[size-leftSize:])
	}
	this.dataSize += size
	return true
}
func (this *sendBuffer) popData() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.dataSize -= this.sendingBlock.size
	this.sendingBlock = nil
}
func (this *sendBuffer) takeSendData() ([]byte, int) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if this.sendingBlock != nil {
		return nil, 0
	}
	if this.blocks.Len() == 0 {
		return nil, 0
	}
	b := this.blocks.Front()
	this.sendingBlock = b.Value.(*block)
	this.blocks.Remove(b)
	return this.sendingBlock.buff, this.sendingBlock.size
}
