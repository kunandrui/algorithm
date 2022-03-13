package depth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func depth(root *TreeNode) int {
	ans := 0
	dep := 0

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			ans = max(ans, dep)
			return
		}
		dep++
		dfs(root.Left)
		dfs(root.Right)
		dep--
	}
	dfs(root)
	return ans
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return max(leftDepth, rightDepth) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
