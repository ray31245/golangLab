package main

import "github.com/emirpasic/gods/stacks/linkedliststack"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	rightPoint := root.Right
	leftPoint := root.Left

	sR := linkedliststack.New()
	sL := linkedliststack.New()

	for {
		if (leftPoint == nil) != (rightPoint == nil) {
			return false
		}
		if (leftPoint != nil && rightPoint != nil) && (rightPoint.Val != leftPoint.Val) {
			return false
		}
		if leftPoint != nil {
			sL.Push(leftPoint)
			sR.Push(rightPoint)
			leftPoint = leftPoint.Left
			if rightPoint == nil {
				return false
			}
			rightPoint = rightPoint.Right
		} else {
			psL, okL := sL.Pop()
			psR, okR := sR.Pop()
			if !(okL && okR) {
				break
			}
			leftPoint = psL.(*TreeNode).Right
			rightPoint = psR.(*TreeNode).Left
		}
	}

	return true
}
