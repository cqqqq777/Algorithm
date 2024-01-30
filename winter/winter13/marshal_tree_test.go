package winter13

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}
	c := Constructor()
	data := c.serialize(root)
	head := c.deserialize(data)
	fmt.Println(head)
}
