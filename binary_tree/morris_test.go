package binary_tree

import "testing"

func TestMorris(t *testing.T) {
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
	MorrisPre(head) // expected output: 1 2 4 5 3 6 7
	MorrisIn(head)  // expected output: 4 2 5 1 6 3 7
	MorrisPos(head) // expected output: 4 5 2 6 7 3 1
}
