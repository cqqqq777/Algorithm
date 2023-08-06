package day06

type treeNode struct {
	v           int
	left, right *treeNode
}

// ConvertTreeToList
// 双向链表节点结构和二叉树节点结构是一样的，如果你把last认为是left，next认为是next的话。
// 给定一个搜索二叉树的头节点head，请转化成一条有序的双向链表，并返回链表的头节点。
func ConvertTreeToList(root *treeNode) *treeNode {
	if root == nil {
		return nil
	}
	head, _ := rec(root)
	return head
}

// 递归
func rec(node *treeNode) (*treeNode, *treeNode) {
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
	var head, tail *treeNode
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
