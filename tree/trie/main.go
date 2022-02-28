package main

type node struct {
	count int
	child []*node
}

type Trie struct {
	root *node
}

func Constructor() Trie {
	t := Trie{
		root: &node{0, make([]*node, 26)},
	}
	return t
}

func (this *Trie) Insert(word string) {
	curr := this.root
	chars := []byte(word)
	for i := 0; i < len(chars); i++ {
		index := int(chars[i] - 'a')
		if curr.child[index] != nil {
			curr = curr.child[index]
		} else {
			curr.child[index] = &node{0, make([]*node, 26)}
			curr = curr.child[index]
		}
	}
	curr.count += 1
}

func (this *Trie) Search(word string) bool {
	curr := this.root
	for i := 0; i < len(word); i++ {
		index := int(word[i] - 'a')
		if curr.child[index] == nil {
			return false
		} else {
			curr = curr.child[index]
		}
	}
	return curr.count > 0
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for i := 0; i < len(prefix); i++ {
		index := int(prefix[i] - 'a')
		if curr.child[index] == nil {
			return false
		} else {
			curr = curr.child[index]
		}
	}
	return true
}

func main() {

}
