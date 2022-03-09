package bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		//插入
		return &TreeNode{val, nil, nil}
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	return findNext(root, p.Val)
}

func findNext(root *TreeNode, p int) *TreeNode {
	var ans *TreeNode
	curr := root
	for curr != nil {
		if p == curr.Val {
			if curr.Right != nil {
				next := curr.Right
				for next.Left != nil {
					next = next.Left
				}
				ans = next
			}
		}
		if p < curr.Val {
			//经过的点找一个最小的
			if ans == nil || ans.Val > curr.Val {
				ans = curr
			}
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return ans
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key == root.Val {
		//删除
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		next := findNext(root, key)
		deleteNode(root, next.Val)
		root.Val = next.Val
		return root
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
