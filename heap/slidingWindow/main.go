package slidingWindow

import (
	"container/heap"
)

type pair struct {
	key   int
	index int
}

type pairHeap []pair

func (h pairHeap) Len() int {
	return len(h)
}

func (h pairHeap) Less(i, j int) bool {
	return h[i].key > h[j].key
}

func (h pairHeap) Swap(i, j int) {
	tmp := h[i]
	h[i] = h[j]
	h[j] = tmp
}

func (h *pairHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *pairHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n]
	*h = (*h)[0 : n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	var ans []int
	h := &pairHeap{}
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		heap.Push(h, pair{nums[i], i})

		if i >= k-1 {
			for (*h)[0].index < i-k+1 {
				heap.Pop(h)
			}
			ans = append(ans, (*h)[0].key)
		}
	}
	return ans
}
