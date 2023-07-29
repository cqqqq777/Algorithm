/*
Package sort_map implements a map with AVL tree.
*/

package sort_map

type AVLNode struct {
	key                 Key
	value               Value
	left, right, parent *AVLNode
	height              int
}

type AVLTree struct {
	root *AVLNode
}

// NewAVLTree return AVL Tree
func NewAVLTree() *AVLTree {
	return new(AVLTree)
}

func (avl *AVLTree) Put(key Key, value Value) {
	if avl.root == nil {
		avl.root = &AVLNode{key: key, value: value}
		return
	}
	avl.put(key, value)
}

func (avl *AVLTree) put(key Key, value Value) {
	cur := avl.root
	for {
		if cur.key > key {
			if cur.left == nil {
				cur.left = &AVLNode{key: key, value: value, parent: cur}
				// update height
				avl.updateHeight(cur)
				avl.reBalance(cur.left)
				return
			}
			cur = cur.left
		} else if cur.key < key {
			if cur.right == nil {
				cur.right = &AVLNode{key: key, value: value, parent: cur}
				// update height
				avl.updateHeight(cur)
				avl.reBalance(cur.right)
				return
			}
			cur = cur.right
		} else {
			// key existed, update the value
			cur.value = value
			return
		}
	}
}

func (avl *AVLTree) updateHeight(node *AVLNode) {
	for parent := node; parent != nil; parent = parent.parent {
		leftHeight, rightHeight := avl.getHeight(parent.left), avl.getHeight(parent.right)
		newHeight := max(leftHeight, rightHeight) + 1
		if newHeight != parent.height {
			parent.height = newHeight
		} else {
			break
		}
	}
}

func (avl *AVLTree) getHeight(node *AVLNode) int {
	if node == nil {
		return -1
	}
	return node.height
}

// Get returns the value of the key.
func (avl *AVLTree) Get(key Key) (Value, bool) {
	return avl.get(avl.root, key)
}

func (avl *AVLTree) get(node *AVLNode, key Key) (Value, bool) {
	cur := node
	for cur != nil {
		if cur.key > key {
			cur = cur.left
		} else if cur.key < key {
			cur = cur.right
		} else {
			return cur.value, true
		}
	}
	return 0, false
}

// Delete deletes the key.
func (avl *AVLTree) Delete(key Key) {
	avl.delete(avl.root, key)
}

func (avl *AVLTree) delete(node *AVLNode, key Key) {
	// search key first
	cur := node
	for cur != nil && cur.key != key {
		if cur.key > key {
			cur = cur.left
		} else if cur.key < key {
			cur = cur.right
		}
	}
	if cur == nil {
		return
	}
	// delete the node
	avl.deleteNode(cur)
}

func (avl *AVLTree) deleteNode(node *AVLNode) {
	// if node has no child
	if node.left == nil && node.right == nil {
		if node.parent == nil {
			avl.root = nil
			return
		}
		if node.parent.left == node {
			node.parent.left = nil
		} else {
			node.parent.right = nil
		}
		// update height
		avl.updateHeight(node.parent)
		avl.reBalance(node.parent)
		return
	}
	// if node has one child
	if node.left == nil {
		if node.parent == nil {
			avl.root = node.right
			avl.root.parent = nil
			return
		}
		if node.parent.left == node {
			node.parent.left = node.right
			node.right.parent = node.parent
		} else {
			node.parent.right = node.right
			node.right.parent = node.parent
		}
		// update height
		avl.updateHeight(node.parent)
		avl.reBalance(node.parent)
		return
	}
	if node.right == nil {
		if node.parent == nil {
			avl.root = node.left
			avl.root.parent = nil
			return
		}
		if node.parent.left == node {
			node.parent.left = node.left
			node.left.parent = node.parent
		} else {
			node.parent.right = node.left
			node.left.parent = node.parent
		}
		// update height
		avl.updateHeight(node.parent)
		avl.reBalance(node.parent)
		return
	}
	// if node has two children
	// find the successor
	successor := node.right
	for successor.left != nil {
		successor = successor.left
	}
	// copy successor's key and value to node
	node.key = successor.key
	node.value = successor.value
	// delete successor
	successor.parent.left = successor.right
	if successor.right != nil {
		successor.right.parent = successor.parent
	}
	// update height
	avl.updateHeight(successor.parent)
	avl.reBalance(successor.parent)
}

// leftRotate rotates the node to left.
func (avl *AVLTree) leftRotate(node *AVLNode) *AVLNode {
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
	avl.updateHeight(node)
	if right.parent != nil {
		if right.parent.left == node {
			right.parent.left = right
		} else {
			right.parent.right = right
		}
	} else {
		avl.root = right
	}
	return right
}

// rightRotate rotates the node to right.
func (avl *AVLTree) rightRotate(node *AVLNode) *AVLNode {
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
	avl.updateHeight(node)
	if left.parent != nil {
		if left.parent.left == node {
			left.parent.left = left
		} else {
			left.parent.right = left
		}
	} else {
		avl.root = left
	}
	return left
}

// reBalance check the balance of the tree and re-balance it.
func (avl *AVLTree) reBalance(node *AVLNode) {
	for ; node != nil; node = node.parent {
		leftHeight, rightHeight := avl.getHeight(node.left), avl.getHeight(node.right)
		if leftHeight-rightHeight == 2 {
			if node.left.left != nil && avl.getHeight(node.left.left) > avl.getHeight(node.left.right) {
				// left-left
				avl.rightRotate(node)
				break
			} else {
				// left-right
				avl.leftRotate(node.left)
				avl.rightRotate(node)
				break
			}
		} else if leftHeight-rightHeight == -2 {
			if avl.getHeight(node.right.right) > avl.getHeight(node.right.left) {
				// right-right
				avl.leftRotate(node)
				break
			} else {
				// right-left
				avl.rightRotate(node.right)
				avl.leftRotate(node)
				break
			}
		}
	}
}
