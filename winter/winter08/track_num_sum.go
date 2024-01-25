package winter08

import "strconv"

// 给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
// 每条从根节点到叶节点的路径都代表一个数字：
// 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
// 计算从根节点到叶节点生成的 所有数字之和 。
// 叶节点 是指没有子节点的节点。
func sumNumbers(root *TreeNode) int {
	var res int
	var back func(node *TreeNode, path string)
	back = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		path += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			sum, _ := strconv.Atoi(path)
			res += sum
			return
		}
		back(node.Left, path)
		back(node.Right, path)
	}
	back(root, "")
	return res
}
