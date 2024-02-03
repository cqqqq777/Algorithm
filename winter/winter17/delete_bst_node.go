package winter17

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。
// 返回二叉搜索树（有可能被更新）的根节点的引用。
func deleteNode(root *TreeNode, key int) *TreeNode {
	// 如果被删除节点没有孩子节点，直接删除
	// 如果被删除节点有一个孩子节点，让被删除节点的父节点接管子树
	// 如果被删除节点有两个孩子节点，找到右子树的最左节点继承他
	if root == nil {
		return root
	}

	dummy := &TreeNode{
		Left:  root,
		Right: root,
	}
	parent, cur := dummy, root
	for cur != nil && cur.Val != key {
		parent = cur
		if cur.Val > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	// 树中没有这个节点
	if cur == nil {
		return root
	}

	// 没有孩子节点
	if cur.Left == nil && cur.Right == nil {
		if parent.Left == cur {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return dummy.Left
	}

	// 有一个孩子节点
	if cur.Right == nil {
		if parent.Left == cur {
			parent.Left = cur.Left
		} else {
			parent.Right = cur.Left
		}
		return dummy.Left
	}
	if cur.Left == nil {
		if parent.Left == cur {
			parent.Left = cur.Right
		} else {
			parent.Right = cur.Right
		}
		return dummy.Left
	}

	// 有两个孩子节点
	successor, sParent := cur.Right, cur
	for successor.Left != nil {
		sParent = successor
		successor = successor.Left
	}
	if successor.Right != nil {
		if sParent.Left == cur {
			sParent.Left = successor.Right
		} else {
			sParent.Right = successor.Right
		}
	}
	cur.Val = successor.Val
	return dummy.Left
}
