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
	Next  unsafe.Pointer
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
	node := &Node{Value: v, Next: nil}
	for {
		tail := load(&lfq.Tail)
		next := load(&tail.Next)
		if tail == load(&lfq.Tail) {
			if next == nil {
				if cas(&tail.Next, next, node) {
					cas(&lfq.Tail, tail, node)
					return
				}
			} else {

			}
		}

	}

}

func load(p *unsafe.Pointer) *Node {
	return (*Node)(atomic.LoadPointer(p))
}

func cas(p *unsafe.Pointer, old, new *Node) bool {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}
