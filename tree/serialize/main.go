package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	seq []string
	cur int
}

func Constructor() Codec {
	c := Codec{
		seq: []string{},
		cur: 0,
	}
	return c
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			this.seq = append(this.seq, "nil")
			return
		}
		val := root.Val
		this.seq = append(this.seq, strconv.Itoa(val))
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return strings.Join(this.seq, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	split := strings.Split(data, ",")
	this.seq = append([]string(nil), split...)
	this.cur = 0
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		str := this.seq[this.cur]
		if str == "nil" {
			this.cur++
			return nil
		}
		val, _ := strconv.Atoi(str)
		root := &TreeNode{val, nil, nil}
		this.cur++
		root.Left = dfs()
		root.Right = dfs()
		return root
	}
	return dfs()
}
