package day06

import "math"

// MaxBSTNodes
// 找到一棵二叉树中，最大的搜索二叉子树，返回最大搜索二叉子树的节点个数。
func MaxBSTNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxBSTNodes(root).maxBSTSize
}

type info struct {
	maxBSTHead           *TreeNode
	max, min, maxBSTSize int
}

func maxBSTNodes(node *TreeNode) *info {
	if node == nil {
		return &info{maxBSTHead: nil, maxBSTSize: 0, min: math.MaxInt, max: math.MinInt}
	}
	lInfo := maxBSTNodes(node.left)
	rInfo := maxBSTNodes(node.right)
	curInfo := &info{max: max(max(lInfo.max, rInfo.max), node.v), min: min(min(lInfo.min, rInfo.min), node.v), maxBSTSize: max(lInfo.maxBSTSize, rInfo.maxBSTSize)}
	if lInfo.maxBSTSize > rInfo.maxBSTSize {
		curInfo.maxBSTHead = lInfo.maxBSTHead
	} else {
		curInfo.maxBSTHead = rInfo.maxBSTHead
	}
	if node.left == lInfo.maxBSTHead && node.right == rInfo.maxBSTHead && node.v > lInfo.max && node.v < rInfo.min {
		curInfo.maxBSTHead = node
		curInfo.maxBSTSize = lInfo.maxBSTSize + rInfo.maxBSTSize + 1
	}
	return curInfo
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
