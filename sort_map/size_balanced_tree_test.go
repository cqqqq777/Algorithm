package sort_map

import (
	"fmt"
	"testing"
)

func TestSBTree(t *testing.T) {
	s := NewSBTree()
	s.Put(1, 1)
	s.Put(2, 2)
	s.Put(3, 3)
	s.Put(4, 4)
	s.Put(5, 5)
	s.Put(6, 6)
	s.Put(7, 7)
	s.Put(8, 8)
	s.Put(9, 9)
	s.Delete(4)
	s.Put(32, 2)
	fmt.Println(s.root.left.key)
	fmt.Println(s.Get(32))
	preOrder(s.root)
}

func preOrder(node *SBTNode) {
	if node == nil {
		return
	}
	preOrder(node.left)
	fmt.Println(node.key)
	preOrder(node.right)
}
