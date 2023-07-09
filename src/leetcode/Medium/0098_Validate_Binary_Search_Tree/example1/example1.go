package example1

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	// MIN, MAX := -1<<63, 1<<63-1
	MIN, MAX := math.MinInt, math.MaxInt
	return dfs(MIN, MAX, root)
}

func dfs(min int, max int, root *TreeNode) bool {
	if root == nil {
		return true
	}
	return min < root.Val && root.Val < max &&
		dfs(min, root.Val, root.Left) &&
		dfs(root.Val, max, root.Right)
}
