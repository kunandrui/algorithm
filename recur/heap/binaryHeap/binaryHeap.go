package binaryHeap

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Key  int
	List *ListNode
}

type BinaryHeap struct {
	Heap []Node
}

func NewBinaryHeap() *BinaryHeap {
	h := BinaryHeap{
		Heap: append([]Node(nil), Node{0, nil}),
	}
	return &h
}

func (b *BinaryHeap) Insert(n Node) {
	b.Heap = append(b.Heap, n)
	b.heapifyUp(len(b.Heap) - 1)
}

func (b *BinaryHeap) RemoveMin() {
	swap(b.Heap, 1, len(b.Heap)-1)
	b.Heap = b.Heap[0 : len(b.Heap)-1]
	b.heapifyDown(1)
}

func (b *BinaryHeap) GetMin() Node {
	return b.Heap[1]
}

func (b *BinaryHeap) Empty() bool {
	return len(b.Heap) == 1
}

func (b *BinaryHeap) heapifyUp(p int) {
	for p > 1 {
		fa := p / 2
		if b.Heap[p].Key < b.Heap[fa].Key {
			swap(b.Heap, p, fa)
			p = fa
		} else {
			break
		}
	}
}

func (b *BinaryHeap) heapifyDown(p int) {
	child := p * 2
	for child < len(b.Heap) {
		another := child + 1
		if another < len(b.Heap) && b.Heap[another].Key < b.Heap[child].Key {
			child = another
		}
		if b.Heap[child].Key < b.Heap[p].Key {
			swap(b.Heap, child, p)
			child = child * 2
		} else {
			break
		}
	}
}

func swap(array []Node, i int, j int) {
	tmp := array[i]
	array[i] = array[j]
	array[j] = tmp
}

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{0, nil}
	cur := head

	h := NewBinaryHeap()

	for _, list := range lists {
		if list != nil {
			h.Insert(Node{Key: list.Val, List: list})
		}
	}

	for !h.Empty() {
		min := h.GetMin()
		cur.Next = &ListNode{min.Key, nil}
		cur = cur.Next

		h.RemoveMin()
		nextListNode := min.List.Next
		if nextListNode != nil {
			nextNode := Node{nextListNode.Val, nextListNode}
			h.Insert(nextNode)
		}
	}
	return head.Next
}
