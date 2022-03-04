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
