package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{root}
	postorder := []int{}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		postorder = append([]int{node.Val}, postorder...)
		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
	}
	return postorder
}
