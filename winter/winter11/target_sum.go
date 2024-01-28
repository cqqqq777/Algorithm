package winter11

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。
// 判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。
// 如果存在，返回 true ；否则，返回 false 。
// 叶子节点 是指没有子节点的节点。
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var back func(node *TreeNode, curSum int) bool
	back = func(node *TreeNode, curSum int) bool {
		if node == nil {
			return false
		}

		curSum += node.Val

		if node.Left == nil && node.Right == nil {
			if curSum == targetSum {
				return true
			}
			return false
		}

		return back(node.Left, curSum) || back(node.Right, curSum)
	}
	return back(root, 0)
}
