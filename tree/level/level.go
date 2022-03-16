package level

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Pair struct {
	Node  *TreeNode
	Depth int
}

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	var deque []Pair
	if root == nil {
		return ans
	}
	deque = append(deque, Pair{root, 0})
	for len(deque) != 0 {
		head := deque[0]
		deque = deque[1:len(deque)]
		depth := head.Depth
		node := head.Node
		if depth >= len(ans) {
			ans = append(ans, []int{})
		}
		ans[depth] = append(ans[depth], node.Val)

		leftChild := node.Left
		if leftChild != nil {
			deque = append(deque, Pair{leftChild, depth + 1})
		}
		rightChild := node.Right
		if rightChild != nil {
			deque = append(deque, Pair{rightChild, depth + 1})
		}
	}

	return ans
}

func levelOrder2(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var deque []*TreeNode
	deque = append(deque, root)
	for len(deque) != 0 {
		sz := len(deque)
		var level []int
		for i := 0; i < sz; i++ {
			node := deque[0]
			deque = deque[1:len(deque)]
			level = append(level, node.Val)
			if node.Left != nil {
				deque = append(deque, node.Left)
			}
			if node.Right != nil {
				deque = append(deque, node.Right)
			}
		}
		ans = append(ans, level)
	}
	return ans
}
