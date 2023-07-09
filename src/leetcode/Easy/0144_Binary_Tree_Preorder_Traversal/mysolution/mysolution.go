package main

import "github.com/emirpasic/gods/stacks/linkedliststack"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraversal(root *TreeNode) []int {
	return preorderTraversal(root)
}

// try to use stack

func preorderTraversal(root *TreeNode) []int {
	s := linkedliststack.New()
	res := []int{}
	walk := root
	for {
		if walk != nil {
			res = append(res, walk.Val)
			s.Push(walk)
			walk = walk.Left
		} else {
			ps, ok := s.Pop()
			if !ok {
				return res
			}
			walk = ps.(*TreeNode).Right
		}
	}
}
