package winter04

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
// 路径和 是路径中各节点值的总和。
// 返回最大路径和
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := root.Val

	var recur func(*TreeNode) int
	recur = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(recur(node.Left), 0)
		right := max(recur(node.Right), 0)

		res = max(res, left+right+node.Val)

		return node.Val + max(left, right)
	}
	recur(root)
	return res
}
