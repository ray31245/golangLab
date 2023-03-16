package main

import "github.com/emirpasic/gods/stacks/linkedliststack"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	return levelOrder(root)
}

func levelOrder(root *TreeNode) [][]int {
	s := linkedliststack.New()
	res := [][]int{}
	walk := root

	for {
		if walk != nil {
			level := s.Size()
			if level >= len(res) {
				res = append(res, []int{})
			}
			res[level] = append(res[level], walk.Val)
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
