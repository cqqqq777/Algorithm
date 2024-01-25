package winter08

// 二叉树最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var res int
	var recur func(node *TreeNode, depth int)
	recur = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		depth++
		res = max(depth, res)
		recur(node.Left, depth)
		recur(node.Right, depth)
	}
	recur(root, 0)
	return res
}
