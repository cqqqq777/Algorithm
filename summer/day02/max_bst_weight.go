package day02

type Node struct {
	weight      int
	left, right *Node
}

// MaxWeight 给定二叉树返回权值最大路径的权值
func MaxWeight(root *Node) int {
	return maxWeight(root)
}

// 暴力递归
func maxWeight(node *Node) int {
	if node == nil {
		return 0
	}
	leftWeight, rightWeight := maxWeight(node.left), maxWeight(node.right)
	if leftWeight > rightWeight {
		return leftWeight + node.weight
	}
	return rightWeight + node.weight
}
