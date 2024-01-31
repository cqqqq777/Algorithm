package winter14

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 判断二叉树是不是完全二叉树
func isCompleteTree(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	leaf := false
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if (n.Left == nil && n.Right != nil) || (leaf && (n.Left != nil || n.Right != nil)) {
			return false
		}
		if n.Left != nil {
			queue = append(queue, n.Left)
		}
		if n.Right != nil {
			queue = append(queue, n.Right)
		} else {
			leaf = true
		}
	}
	return true
}
