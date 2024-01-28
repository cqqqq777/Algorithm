package winter11

// 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
func invertTree(root *TreeNode) *TreeNode {
	// 左右节点交换即可
	var recur func(node *TreeNode)
	recur = func(node *TreeNode) {
		if node == nil {
			return
		}
		node.Left, node.Right = node.Right, node.Left
		recur(node.Left)
		recur(node.Right)
	}
	recur(root)
	return root
}
