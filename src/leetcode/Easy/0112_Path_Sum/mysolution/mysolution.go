package mysolution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func HasPathSum(root *TreeNode, targetSum int) bool {
	return hasPathSum(root, targetSum)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	w := walk{targetSum: targetSum}
	return w.dfs(root, 0)
}

type walk struct {
	targetSum int
}

func (w walk) dfs(node *TreeNode, sum int) bool {
	if node == nil {
		return false
	}

	sum = node.Val + sum
	if sum == w.targetSum && (node.Left == nil && node.Right == nil) {
		return true
	}

	l := w.dfs(node.Left, sum)
	if l {
		return true
	}

	r := w.dfs(node.Right, sum)
	if r {
		return true
	}

	return false
}
