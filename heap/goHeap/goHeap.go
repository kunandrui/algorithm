package main

import (
	"container/heap"
)

type Node struct {
	Key int
}

type NodeHeap []Node

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].Key > h[j].Key
}

func (h NodeHeap) Swap(i, j int) {
	tmp := h[i]
	h[i] = h[j]
	h[j] = tmp
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *NodeHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func main() {
	h := &NodeHeap{}
	for i := 0; i < 10; i++ {
		*h = append(*h, Node{i})
	}

	heap.Init(h)
	heap.Push(h, Node{0})
	heap.Pop(h)
}
