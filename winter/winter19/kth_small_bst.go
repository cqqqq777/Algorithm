package winter19

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
func kthSmallest(root *TreeNode, k int) int {
	var res int
	var recur func(node *TreeNode, k *int)
	recur = func(node *TreeNode, k *int) {
		if node == nil {
			return
		}
		recur(node.Left, k)
		*k--
		if *k == 0 {
			res = node.Val
		}
		recur(node.Right, k)
	}
	recur(root, &k)
	return res
}
