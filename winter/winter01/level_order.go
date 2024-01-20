package winter01

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树层序遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		sz := len(queue)
		curLevel := make([]int, 0, sz)
		for i := 0; i < sz; i++ {
			node := queue[0]
			queue = queue[1:]
			curLevel = append(curLevel, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, curLevel)
	}
	return res
}
