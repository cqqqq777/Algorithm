package day06

type TreeNode struct {
	v           int
	left, right *TreeNode
}

// ConvertTreeToList
// 双向链表节点结构和二叉树节点结构是一样的，如果你把last认为是left，next认为是next的话。
// 给定一个搜索二叉树的头节点head，请转化成一条有序的双向链表，并返回链表的头节点。
func ConvertTreeToList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	head, _ := rec(root)
	return head
}

// 递归
func rec(node *TreeNode) (*TreeNode, *TreeNode) {
	if node == nil {
		return nil, nil
	}
	lHead, lTail := rec(node.left)
	rHead, rTail := rec(node.right)
	if lTail != nil {
		lTail.right = node
	}
	node.left = lTail
	node.right = rHead
	if rHead != nil {
		rHead.left = node
	}
	var head, tail *TreeNode
	if lHead == nil {
		head = node
	} else {
		head = lHead
	}
	if rTail == nil {
		tail = node
	} else {
		tail = rTail
	}
	return head, tail
}
