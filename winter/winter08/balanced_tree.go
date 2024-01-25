package winter08

// 给定一个二叉树，判断它是否是高度平衡的二叉树。
// 本题中，一棵高度平衡二叉树定义为：
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var recur func(node *TreeNode) (bool, int)
	recur = func(node *TreeNode) (bool, int) {
		if node == nil {
			return true, 0
		}
		lOk, lDepth := recur(node.Left)
		rOk, rDepth := recur(node.Right)
		depth := max(lDepth, rDepth) + 1
		if !lOk || !rOk {
			return false, depth
		}
		diff := lDepth - rDepth
		if diff > 1 || diff < -1 {
			return false, depth
		}
		return true, depth
	}
	f, _ := recur(root)
	return f
}
