package LRU

type Node struct {
	key, value int
	next, prev *Node
}

type LRUCache struct {
	capacity   int
	head, tail *Node
	hash       map[int]*Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		head:     &Node{0, 0, nil, nil},
		tail:     &Node{0, 0, nil, nil},
		hash:     map[int]*Node{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if v, exist := this.hash[key]; exist {
		remove(v)
		insert(this.head, v)
		return v.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, exist := this.hash[key]; exist {
		v.value = value
		remove(v)
		insert(this.head, v)

	} else {
		node := &Node{key, value, nil, nil}
		this.hash[key] = node
		insert(this.head, node)
		if len(this.hash) > this.capacity {
			delete(this.hash, this.tail.prev.key)
			remove(this.tail.prev)
		}
	}
}

func remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func insert(p *Node, node *Node) {
	//p.next和node
	p.next.prev = node
	node.next = p.next

	//p和node
	node.prev = p
	p.next = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
