package winter20

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// recur 返回每棵每棵子树转换为链表后的头和尾
	var recur func(node *TreeNode) (*TreeNode, *TreeNode)
	recur = func(node *TreeNode) (*TreeNode, *TreeNode) {
		if node == nil {
			return nil, nil
		}

		lHead, lTail := recur(node.Left)
		rHead, rTail := recur(node.Right)
		node.Left = nil

		if lHead != nil && rHead != nil {
			node.Right = lHead
			lTail.Right = rHead
			return node, rTail
		}

		if lHead != nil {
			node.Right = lHead
			return node, lTail
		}

		if rHead != nil {
			node.Right = rHead
			return node, rTail
		}

		return node, node
	}
	recur(root)
	return
}
