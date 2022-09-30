package service

import (
	"sync"
	"sync/atomic"

	// "time"
	"unsafe"
)

const (
	MAX_DATA_SIZE = 100
)

// lock free queue
type Queue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
	sema uint32
	cond sync.Cond
}

// one node in queue
type Node struct {
	cmd  ICommand
	next unsafe.Pointer
}

// queue functions
func (self *Queue) enQueue(cmd ICommand) {
	newValue := unsafe.Pointer(&Node{cmd: cmd, next: nil})
	var tail, next unsafe.Pointer
	for {
		tail = self.tail
		next = ((*Node)(tail)).next
		if next != nil {
			atomic.CompareAndSwapPointer(&(self.tail), tail, next)
		} else if atomic.CompareAndSwapPointer(&((*Node)(tail).next), nil, newValue) {
			self.cond.L.Lock()
			self.cond.Signal()
			self.cond.L.Unlock()
			break
		}
	}
}

func (self *Queue) deQueue() (cmd ICommand) {
	var head, tail, next unsafe.Pointer
	for {
		head = self.head
		tail = self.tail
		next = ((*Node)(head)).next
		if head == tail {
			if next == nil {
				self.cond.Wait()
			} else {
				atomic.CompareAndSwapPointer(&(self.tail), tail, next)
			}
		} else {
			cmd = ((*Node)(next)).cmd
			if atomic.CompareAndSwapPointer(&(self.head), head, next) {
				return cmd
			}
		}
	}
}

func makeQueue() *Queue {
	queue := new(Queue)
	queue.head = unsafe.Pointer(new(Node))
	queue.tail = queue.head
	var mutex sync.Mutex
	queue.cond = sync.Cond{L: &mutex}
	return queue
}
