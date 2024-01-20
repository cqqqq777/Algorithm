package winter03

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	// 层序遍历使用一个遍历记录正序还是逆序即可
	var flag bool
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		sz := len(queue)
		level := make([]int, sz)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			if flag {
				level[sz-i-1] = cur.Val
			} else {
				level[i] = cur.Val
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, level)
		flag = !flag
	}
	return res
}
