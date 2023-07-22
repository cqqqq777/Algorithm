package binary_tree

import "math"

type ReturnInfo struct {
	maxDistance, height int
}

// MaxDistance 返回二叉树的两节点间的最大距离
func MaxDistance(head *Node) int {
	return recursion(head).maxDistance
}

func recursion(node *Node) *ReturnInfo {
	info := new(ReturnInfo)
	if node == nil {
		return info
	}
	leftInfo, rightInfo := recursion(node.Left), recursion(node.Right)

	d1, d2 := leftInfo.maxDistance, rightInfo.maxDistance
	d3 := leftInfo.height + rightInfo.height + 1
	info.maxDistance = int(math.Max(float64(d1), math.Max(float64(d2), float64(d3))))
	if leftInfo.height > rightInfo.height {
		info.height = leftInfo.height + 1
	} else {
		info.height = rightInfo.height + 1
	}
	return info
}
