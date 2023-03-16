package main

import "github.com/emirpasic/gods/stacks/linkedliststack"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	return inorderTraversal(root)
}

// try to use stack

func inorderTraversal(root *TreeNode) []int {
	s := linkedliststack.New()
	res := []int{}
	walk := root
	for {
		if walk != nil {
			s.Push(walk)
			walk = walk.Left
		} else {
			ps, ok := s.Pop()
			if !ok {
				return res
			}
			res = append(res, ps.(*TreeNode).Val)
			walk = ps.(*TreeNode).Right
		}
	}
}
