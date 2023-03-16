package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PostorderTraversal(root *TreeNode) []int {
	return postorderTraversal(root)
}

func postorderTraversal(root *TreeNode) []int {
	var res []int
	postorder(root, &res)
	return res
}

func postorder(root *TreeNode, out *[]int) {
	if root != nil {
		postorder(root.Left, out)
		postorder(root.Right, out)
		*out = append(*out, root.Val)
	}
}
