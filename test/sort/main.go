package main

import (
	"fmt"
	"sort"
)

type node struct {
	a int
	b int
}

func main() {
	nodes := []node{node{1, 1}, node{2, 2}, node{2, 3}, node{2, -1}}

	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].a == nodes[j].a {
			return nodes[i].b < nodes[j].b
		} else {
			return nodes[i].a < nodes[j].a
		}
	})

	fmt.Println(nodes)
}
