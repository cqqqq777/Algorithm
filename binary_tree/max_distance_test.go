package binary_tree

import (
	"fmt"
	"testing"
)

func TestMaxDistance(t *testing.T) {
	head := new(Node)
	head.Data = 1
	head.Left = new(Node)
	head.Left.Data = 2
	head.Right = new(Node)
	head.Right.Data = 3
	head.Left.Left = new(Node)
	head.Left.Left.Data = 4
	head.Left.Right = new(Node)
	head.Left.Right.Data = 5
	head.Right.Left = new(Node)
	head.Right.Left.Data = 6
	head.Right.Right = new(Node)
	head.Right.Right.Data = 7
	/*
							1
				       2         3
		            4     5    6    7
	*/
	fmt.Println(MaxDistance(head))
}
