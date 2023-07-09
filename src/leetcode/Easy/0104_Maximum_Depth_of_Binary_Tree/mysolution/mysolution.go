package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	if l > r {
		return l + 1
	}
	return r + 1
}

// func dfs(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	l := dfs(root.Left) + 1
// 	r := dfs(root.Right) + 1

// 	if l > r {
// 		return l
// 	} else {
// 		return r
// 	}
// }
