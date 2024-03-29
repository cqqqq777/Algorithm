package winter19

// 给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。
// 如果存在，返回 true ；否则，返回 false 。
// 二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	}
	return helper(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func helper(root, subRoot *TreeNode) bool {
	if subRoot == nil && root == nil {
		return true
	}
	if subRoot == nil {
		return false
	}
	if subRoot.Val != root.Val {
		return false
	}
	return helper(root.Left, subRoot.Left) && helper(root.Right, subRoot.Right)
}
