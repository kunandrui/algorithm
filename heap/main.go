package main

import "fmt"

func main() {
	h := NewHeap()
	arr := []int{3, 2, 1, 5, 6, 8, 4, 9, 7, 10}
	for i := 1; i <= 10; i++ {
		h.Insert(arr[i-1])
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(h.Top())
		h.Pop()
	}
}

type Heap struct {
	arr []int
}

func NewHeap() *Heap {
	return &Heap{[]int{0}}
}

func (h *Heap) Empty() bool {
	return len(h.arr) == 1
}

func (h *Heap) Top() int {
	return h.arr[1]
}

func (h *Heap) Insert(val int) {
	h.arr = append(h.arr, val)
	p := len(h.arr) - 1
	h.heapifyUp(p)
}

func (h *Heap) Pop() {
	n := len(h.arr) - 1
	h.swap(1, n)
	//删除最后一个元素
	h.arr = h.arr[:n]
	h.heapifyDown(1)
}

func (h *Heap) heapifyUp(p int) {
	for p > 1 {
		if h.arr[p] < h.arr[p/2] {
			h.swap(p, p/2)
			p = p / 2
		} else {
			break
		}
	}
}

func (h *Heap) heapifyDown(p int) {
	child := p * 2
	for child < len(h.arr) {
		other := child + 1
		if other < len(h.arr) && h.arr[other] < h.arr[child] {
			child = other
		}
		if h.arr[child] < h.arr[p] {
			h.swap(child, p)
			p = child
			child = p * 2
		} else {
			break
		}
	}
}

func (h *Heap) swap(i, j int) {
	tmp := h.arr[i]
	h.arr[i] = h.arr[j]
	h.arr[j] = tmp
}
