package lockfreequeue

import (
	"sync/atomic"
	"unsafe"
)

type LockFreeQueue struct {
	Head unsafe.Pointer
	Tail unsafe.Pointer
}

type Node struct {
	Value interface{}
	Next *Node
}

func NewLockFreeQueue() *LockFreeQueue {
	n := unsafe.Pointer(&Node{})
	lfq := LockFreeQueue{
		Head: n,
		Tail: n,
	}
	return &lfq
}

func (lfq *LockFreeQueue) Enqueue(v interface{}) {
	node := &Node{Value:v, Next: nil}
	for {
		tail := (*Node)(atomic.LoadPointer(&lfq.Tail))
		next := (*Node)(atomic.LoadPointer()
	}

}