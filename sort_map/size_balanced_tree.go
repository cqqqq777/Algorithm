/*
Package sort_map provides a map with Size Balanced Tree.
*/

package sort_map

type SBTNode struct {
	key                 Key
	value               Value
	left, right, parent *SBTNode
	size                int
}

type SBTree struct {
	root *SBTNode
}

func NewSBTree() *SBTree {
	return &SBTree{}
}

func (sbt *SBTree) Get(key Key) (Value, bool) {
	cur := sbt.root
	for cur != nil && cur.key != key {
		if cur.key > key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	if cur == nil {
		return 0, false
	}
	return cur.value, true
}

func (sbt *SBTree) leftRotate(node *SBTNode) *SBTNode {
	if node == nil || node.right == nil {
		return node
	}
	right := node.right
	right.parent = node.parent

	node.right = right.left
	if node.right != nil {
		node.right.parent = node
	}
	right.left = node
	node.parent = right
	if right.parent != nil {
		if right.parent.left == node {
			right.parent.left = right
		} else {
			right.parent.right = right
		}
	} else {
		sbt.root = right
	}
	sbt.updateSize(node)
	return right
}

func (sbt *SBTree) rightRotate(node *SBTNode) *SBTNode {
	if node == nil || node.left == nil {
		return node
	}
	left := node.left
	left.parent = node.parent

	node.left = left.right
	if node.left != nil {
		node.left.parent = node
	}
	left.right = node
	node.parent = left
	if left.parent != nil {
		if left.parent.left == node {
			left.parent.left = left
		} else {
			left.parent.right = left
		}
	} else {
		sbt.root = left
	}
	sbt.updateSize(node)
	return left
}

func (sbt *SBTree) reBalance(node *SBTNode) {
	for parent := node; parent != nil; parent = parent.parent {
		llSize, lrSize, rlSize, rrSize := 0, 0, 0, 0
		if parent.left != nil {
			llSize = sbt.getSize(parent.left.left)
			lrSize = sbt.getSize(parent.left.right)
		}
		if parent.right != nil {
			rlSize = sbt.getSize(parent.right.left)
			rrSize = sbt.getSize(parent.right.right)
		}

		if llSize > sbt.getSize(parent.right) {
			cur := sbt.rightRotate(parent)
			sbt.reBalance(parent)
			sbt.reBalance(cur)
		}
		if lrSize > sbt.getSize(parent.right) {
			sbt.leftRotate(parent.left)
			cur := sbt.rightRotate(parent)
			sbt.reBalance(parent)
			sbt.reBalance(cur.left)
			sbt.reBalance(cur)
		}
		if rrSize > sbt.getSize(parent.left) {
			cur := sbt.leftRotate(parent)
			sbt.reBalance(parent)
			sbt.reBalance(cur)
		}
		if rlSize > sbt.getSize(parent.left) {
			sbt.rightRotate(parent.right)
			cur := sbt.leftRotate(parent)
			sbt.reBalance(parent)
			sbt.reBalance(cur.right)
			sbt.reBalance(cur)
		}
	}
}

// updateSize update the size of the node and its ancestors.
func (sbt *SBTree) updateSize(node *SBTNode) {
	cur := node
	for cur != nil {
		cur.size = sbt.getSize(cur.left) + sbt.getSize(cur.right) + 1
		cur = cur.parent
	}
}

func (sbt *SBTree) getSize(node *SBTNode) int {
	if node == nil {
		return 0
	}
	return node.size
}

func (sbt *SBTree) Put(key Key, value Value) {
	if sbt.root == nil {
		sbt.root = &SBTNode{key: key, value: value, size: 1}
		return
	}
	sbt.put(key, value)
}

func (sbt *SBTree) put(key Key, value Value) {
	cur := sbt.root
	for {
		if cur.key > key {
			if cur.left == nil {
				cur.left = &SBTNode{key: key, value: value, parent: cur, size: 1}
				sbt.updateSize(cur)
				sbt.reBalance(cur)
				return
			}
			cur = cur.left
		} else if cur.key < key {
			if cur.right == nil {
				cur.right = &SBTNode{key: key, value: value, parent: cur, size: 1}
				sbt.updateSize(cur)
				sbt.reBalance(cur)
				return
			}
			cur = cur.right
		} else {
			cur.value = value
			return
		}
	}
}

func (sbt *SBTree) Delete(key Key) {
	sbt.delete(key)
}

func (sbt *SBTree) delete(key Key) {
	cur := sbt.root
	for cur != nil && cur.key != key {
		if cur.key > key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	if cur == nil {
		return
	}
	if cur.left == nil && cur.right == nil {
		if cur.parent == nil {
			sbt.root = nil
			return
		}
		if cur.parent.left == cur {
			cur.parent.left = nil
		} else {
			cur.parent.right = nil
		}
		sbt.updateSize(cur.parent)
		sbt.reBalance(cur.parent)
	}
	if cur.left == nil {
		if cur.parent == nil {
			sbt.root = cur.right
			return
		}
		if cur.parent.left == cur {
			cur.parent.left = cur.right
		} else {
			cur.parent.right = cur.right
		}
		cur.right.parent = cur.parent
		sbt.updateSize(cur.parent)
		sbt.reBalance(cur.parent)
	}
	if cur.right == nil {
		if cur.parent == nil {
			sbt.root = cur.left
			return
		}
		if cur.parent.left == cur {
			cur.parent.left = cur.left
		} else {
			cur.parent.right = cur.left
		}
		cur.left.parent = cur.parent
		sbt.updateSize(cur.parent)
		sbt.reBalance(cur.parent)
	}
	// find the successor
	successor := cur.right
	for successor.left != nil {
		successor = successor.left
	}
	cur.key = successor.key
	cur.value = successor.value
	// delete successor
	successor.parent.left = successor.right
	if successor.right != nil {
		successor.right.parent = successor.parent
	}
	// update height
	sbt.updateSize(successor.parent)
	sbt.reBalance(successor.parent)
}
