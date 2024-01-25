package winter08

// 前序遍历
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var recur func(node *TreeNode)
	recur = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		recur(node.Left)
		recur(node.Right)
	}
	recur(root)
	return res
}
