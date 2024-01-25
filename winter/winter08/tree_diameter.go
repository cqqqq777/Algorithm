package winter08

// 给你一棵二叉树的根节点，返回该树的 直径 。
// 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
// 两节点之间路径的 长度 由它们之间边数表示。
func diameterOfBinaryTree(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 0
	}

	var res int
	var recur func(node *TreeNode) int
	recur = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lDepth := recur(node.Left)
		rDepth := recur(node.Right)
		// 当前子树的直径等于左子树高度加右子树高度
		res = max(lDepth+rDepth, res)
		// 返回子树高度
		return max(lDepth+1, rDepth+1)
	}
	recur(root)
	return res
}
