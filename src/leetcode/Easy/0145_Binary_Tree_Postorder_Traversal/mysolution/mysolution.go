package main

import "github.com/emirpasic/gods/stacks/linkedliststack"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := linkedliststack.New()
	stack.Push(root)
	for !stack.Empty() {
		node, ok := stack.Pop()
		if node.(*TreeNode) == nil || !ok {
			continue
		}
		res = append([]int{node.(*TreeNode).Val}, res...)
		stack.Push(node.(*TreeNode).Left)
		stack.Push(node.(*TreeNode).Right)
	}
	return res
}
