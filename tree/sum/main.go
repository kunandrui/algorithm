package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	var ans [][]int
	var set []int

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, target int) {
		if root == nil {
			return
		}
		target -= root.Val
		set = append(set, root.Val)
		if root.Left == nil && root.Right == nil && target == 0 {
			ans = append(ans, append([]int(nil), set...))
		}
		dfs(root.Left, target)
		dfs(root.Right, target)
		set = set[:len(set)-1]
	}
	dfs(root, target)
	return ans
}
