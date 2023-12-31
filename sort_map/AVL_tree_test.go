package sort_map

import (
	"fmt"
	"testing"
)

func TestAVL(t *testing.T) {
	m := NewAVLTree()
	m.Put(1, 1)
	m.Put(2, 2)
	m.Put(3, 3)
	m.Put(9, 9)
	m.Put(8, 8)
	m.Put(4, 4)
	fmt.Println(m.root.right.left.key)
	m.inOrder(m.root)
}

func (avl *AVLTree) inOrder(node *AVLNode) {
	if node == nil {
		return
	}
	avl.inOrder(node.left)
	fmt.Println(node.key)
	avl.inOrder(node.right)
}
