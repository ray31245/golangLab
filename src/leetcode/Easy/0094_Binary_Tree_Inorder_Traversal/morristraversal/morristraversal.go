package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	return inorderTraversal(root)
}

// threaded Binary tree

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	walk := root
	var prev *TreeNode
	for walk != nil {
		if walk.Left == nil {
			res = append(res, walk.Val)
			walk = walk.Right
		} else {
			prev = walk.Left
			for prev.Right != nil {
				prev = prev.Right
			}
			prev.Right = walk
			temp := walk
			walk = walk.Left
			temp.Left = nil
		}
	}

	return res
}
