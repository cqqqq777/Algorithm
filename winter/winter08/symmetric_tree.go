package winter08

// 给你一个二叉树的根节点 root ， 检查它是否轴对称。
func isSymmetric(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, []*TreeNode{root, root}...)
	for len(queue) != 0 {
		node1, node2 := queue[0], queue[1]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}
		queue = append(queue, []*TreeNode{node1.Left, node2.Right, node1.Right, node2.Left}...)
	}
	return true
}
