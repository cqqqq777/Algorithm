package winter08

import "math"

// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
// 有效 二叉搜索树定义如下：
// 节点的左子树只包含 小于 当前节点的数。
// 节点的右子树只包含 大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
func isValidBST(root *TreeNode) bool {
	var recur func(node *TreeNode) (bool, int, int)

	recur = func(node *TreeNode) (bool, int, int) {
		if node == nil {
			return true, math.MaxInt, math.MinInt
		}
		lIs, lMin, lMax := recur(node.Left)
		rIS, rMin, rMax := recur(node.Right)
		if !lIs || !rIS {
			return false, 0, 0
		}
		flag := true
		if node.Val <= lMax || node.Val >= rMin {
			flag = false
		}
		return flag, min(node.Val, lMin, rMin), max(node.Val, lMax, rMax)
	}

	f, _, _ := recur(root)
	return f
}
