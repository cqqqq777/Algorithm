package day07

type node struct {
	left, right *node
}

// CBTNodes 求完全二叉树的节点个数
func CBTNodes(root *node) int {
	if root == nil {
		return 0
	}
	height := 1
	for cur := root.left; cur != nil; cur = cur.left {
		height++
	}
	return re(root, height)
}

func re(n *node, height int) int {
	if n == nil {
		return 0
	}
	h := 1
	for cur := n.right; cur != nil; cur = cur.left {
		h++
	}
	if h == height {
		return pow(2, height-1) + re(n.right, height-1)
	} else {
		return pow(2, height-2) + re(n.left, height-1)
	}
}

func pow(a, b int) int {
	res := 1
	for i := 0; i < b; i++ {
		res *= a
	}
	return res
}
