package BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func serch(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val == root.Val {
		return root
	}
	var res *TreeNode
	if val < root.Val {
		res = serch(root.Left, val)
	} else {
		res = serch(root.Right, val)
	}
	return res
}

func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		// 插入
		return &TreeNode{val, nil, nil}
	}
	if val < root.Val {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}
	return root
}

// 后继者
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var ans *TreeNode
	curr := root
	for curr != nil {
		if p.Val == curr.Val {
			if curr.Right != nil {
				curr = curr.Right
				for curr.Left != nil {
					curr = curr.Left
				}
				ans = curr
			}
			break
		}
		if p.Val < curr.Val {
			if ans == nil || ans.Val > curr.Val {
				ans = curr
			}
			curr = curr.Left
		}
		if p.Val > curr.Val {
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
		//delete
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		next := root.Right
		for next.Left != nil {
			next = next.Left
		}
		root.Right = deleteNode(root.Right, next.Val)
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
