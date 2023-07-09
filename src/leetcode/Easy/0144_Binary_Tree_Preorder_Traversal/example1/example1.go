package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		res = append(res, root.Val)
		tmp := preorderTraversal(root.Left)
		res = append(res, tmp...)
		tmp = preorderTraversal(root.Right)
		res = append(res, tmp...)
	}
	return res
}
