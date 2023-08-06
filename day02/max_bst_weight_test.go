package day02

import (
	"fmt"
	"testing"
)

func TestMaxWeight(t *testing.T) {
	head := new(Node)
	head.weight = 1
	head.left = new(Node)
	head.left.weight = 2
	head.right = new(Node)
	head.right.weight = 3
	head.left.left = new(Node)
	head.left.left.weight = 4
	head.left.right = new(Node)
	head.left.right.weight = 5
	head.right.left = new(Node)
	head.right.left.weight = 6
	head.right.right = new(Node)
	head.right.right.weight = 7
	/*
							1
				       2         3
		            4     5    6    7
	*/
	fmt.Println(MaxWeight(head)) // expected output: 11
}
