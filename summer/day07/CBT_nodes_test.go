package day07

import (
	"fmt"
	"testing"
)

func TestCBTNodes(t *testing.T) {
	root := new(node)
	root.left = new(node)
	root.right = new(node)
	root.left.left = new(node)
	root.left.right = new(node)
	root.right.left = new(node)
	root.right.right = new(node)
	root.left.left.left = new(node)
	root.left.left.right = new(node)
	root.left.right.left = new(node)
	root.left.right.right = new(node)
	fmt.Println(CBTNodes(root)) // expected output: 11
}
