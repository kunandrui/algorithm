package lowestCommonAncestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode
	var recur func(root, p, q *TreeNode) (bool, bool)
	recur = func(root, p, q *TreeNode) (bool, bool) {
		if root == nil {
			return false, false
		}
		leftHasP, leftHasQ := recur(root.Left, p, q)
		rightHasP, rightHasQ := recur(root.Right, p, q)
		hasP := leftHasP || rightHasP || root.Val == p.Val
		hasQ := leftHasQ || rightHasQ || root.Val == q.Val
		if hasP && hasQ && ans == nil {
			ans = root
		}
		return hasP, hasQ
	}
	recur(root, p, q)
	return ans
}
