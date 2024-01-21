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
	res := node.weight

	var recur func(node2 *Node) int
	recur = func(curNode *Node) int {
		if curNode == nil {
			return 0
		}
		left := max(recur(curNode.left), 0)
		right := max(recur(curNode.right), 0)

		res = max(res, left+right+curNode.weight)

		return curNode.weight + max(left, right)
	}
	recur(node)
	return res
}
