package winter10

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	var track []int
	var back func(node *TreeNode, curSum int)
	back = func(node *TreeNode, curSum int) {
		if node == nil {
			return
		}

		curSum += node.Val
		track = append(track, node.Val)

		if node.Left == nil && node.Right == nil {
			if curSum == targetSum {
				res = append(res, append([]int(nil), track...))
			}
			track = track[:len(track)-1]
			return
		}

		back(node.Left, curSum)
		back(node.Right, curSum)
		track = track[:len(track)-1]
	}

	back(root, 0)
	return res
}
