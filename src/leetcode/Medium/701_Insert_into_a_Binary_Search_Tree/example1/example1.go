package example1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

func insert(n *TreeNode, v int) *TreeNode {
	if n == nil {
		return &TreeNode{Val: v}
	}

	if n.Val < v {
		n.Right = insert(n.Right, v)
	} else {
		n.Left = insert(n.Left, v)
	}

	return n
}
