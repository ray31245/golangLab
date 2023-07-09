package mysolution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return dfs(root, root.Right, root.Left)
}

func dfs(node *TreeNode, max *TreeNode, min *TreeNode) bool {
	if node.Left != nil {
		if min.Val <= node.Left.Val {
			return false
		}
		min = node.Left
		if !dfs(node.Left, node, min) {
			return false
		}
	}
	if node.Right != nil {
		if max.Val >= node.Right.Val {
			return false
		}
		max = node.Right
		if !dfs(node.Right, max, node) {
			return false
		}
	}
	return true
}
