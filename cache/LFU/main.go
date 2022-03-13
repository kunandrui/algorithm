package LFU

type ListNode struct {
	Key  int
	Val  int
	Freq int
	Prev *ListNode
	Next *ListNode
}

type LFU struct {
	Capacity int
	Size     int
	FreqHash map[int]*ListNode
	KeyHash  map[int]*ListNode
}

func NewLFU(cap int) *LFU {
	lfu := &LFU{
		Capacity: cap,
		Size:     0,
		FreqHash: make(map[int]*ListNode),
		KeyHash:  make(map[int]*ListNode),
	}
	return lfu
}

func (l *LFU) get(key int) int {
	node, exists := l.KeyHash[key]
	if !exists {
		return -1
	}

	freq := node.Freq
	//1. 移除
	remove(node)
	freq += 1
	//2. 头部插入
	if l.FreqHash[freq] != nil {
		head := &ListNode{-1, -1, -1, nil, nil}
		tail := &ListNode{-1, -1, -1, nil, nil}
		head.Next = tail
		tail.Prev = head
		l.FreqHash[freq] = head
	}
	insert(l.FreqHash[freq], node)
	return node.Val
}

func (l *LFU) put(key, val int) {
	node, exists := l.KeyHash[key]
	if !exists {
		//1. 插入
		if l.Size == l.Capacity {
			//lfu
			lf := 1
			lfuNode := l.FreqHash[lf].Next
			remove(lfuNode)
			lfuKey := lfuNode.Key
			//这里需要修改
			delete(l.KeyHash, lfuKey)
			l.Size--
		}
		freq := 1
		newNode := &ListNode{
			Key:  key,
			Val:  val,
			Freq: freq,
			Prev: nil,
			Next: nil,
		}
		//头部插入
		if l.FreqHash[freq] != nil {
			head := &ListNode{-1, -1, -1, nil, nil}
			tail := &ListNode{-1, -1, -1, nil, nil}
			head.Next = tail
			tail.Prev = head
			l.FreqHash[freq] = head
		}
		insert(l.FreqHash[freq], newNode)
		l.KeyHash[key] = newNode
		l.Size++
	} else {
		//2. 修改
		freq := node.Freq
		freq++
		node.Freq = freq
		remove(node)
		//头部插入
		if l.FreqHash[freq] != nil {
			head := &ListNode{-1, -1, -1, nil, nil}
			tail := &ListNode{-1, -1, -1, nil, nil}
			head.Next = tail
			tail.Prev = head
			l.FreqHash[freq] = head
		}
		insert(l.FreqHash[freq], node)
	}
}

func remove(node *ListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func insert(p *ListNode, node *ListNode) {
	p.Next.Prev = node
	node.Next = p.Next

	p.Next = node
	node.Prev = p
}
