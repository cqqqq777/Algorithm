package winter05

//  Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的右视图
// 层序遍历，每层的最后一个加入 res 中
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue, res := make([]*TreeNode, 0), make([]int, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			n := queue[0]
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
			if i == sz-1 {
				res = append(res, n.Val)
			}
			queue = queue[1:]
		}
	}
	return res
}
