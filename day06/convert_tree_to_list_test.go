package day06

import (
	"fmt"
	"testing"
)

func TestConvertTreeToList(t *testing.T) {
	head := new(treeNode)
	head.v = 1
	head.left = new(treeNode)
	head.left.v = 2
	head.right = new(treeNode)
	head.right.v = 3
	head.left.left = new(treeNode)
	head.left.left.v = 4
	head.left.right = new(treeNode)
	head.left.right.v = 5
	//head.right.left = new(treeNode)
	//head.right.left.v = 6
	head.right.right = new(treeNode)
	head.right.right.v = 7
	/*
							1
				       2         3
		            4     5         7
	*/
	head = ConvertTreeToList(head)
	printList(head) // expected output: 4  2  5  1   3  7
}

func printList(head *treeNode) {
	cur := head
	for ; cur != nil; cur = cur.right {
		fmt.Printf("%d  ", cur.v)
	}
}
