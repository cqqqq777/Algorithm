package day06

import (
	"fmt"
	"testing"
)

func TestMaxBSTNodes(t *testing.T) {
	head := new(treeNode)
	head.v = 1
	head.left = new(treeNode)
	head.left.v = 2
	head.right = new(treeNode)
	head.right.v = 3
	head.left.left = new(treeNode)
	head.left.left.v = 4
	head.left.left.left = new(treeNode)
	head.left.left.left.v = 3
	head.left.left.right = new(treeNode)
	head.left.left.right.v = 5
	head.left.right = new(treeNode)
	head.left.right.v = 5
	head.right.left = new(treeNode)
	head.right.left.v = 6
	head.right.right = new(treeNode)
	head.right.right.v = 7
	/*
								1
					       2         3
			            4     5   6     7
		              3   5
	*/
	fmt.Println(MaxBSTNodes(head)) // expected output: 3
}
