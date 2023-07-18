package binary_tree

import (
	"fmt"
	"testing"
)

func TestIsBST(t *testing.T) {
	head := new(Node)
	head.Data = 10
	head.Left = new(Node)
	head.Left.Data = 6
	head.Right = new(Node)
	head.Right.Data = 16
	head.Left.Left = new(Node)
	head.Left.Left.Data = 3
	head.Left.Right = new(Node)
	head.Left.Right.Data = 9
	head.Right.Left = new(Node)
	head.Right.Left.Data = 14
	head.Right.Right = new(Node)
	head.Right.Right.Data = 19
	fmt.Println(IsBST(head))
	fmt.Println(IsBST2(head))
}

func TestIsFBT(t *testing.T) {
	head := new(Node)
	head.Left = new(Node)
	head.Right = new(Node)
	head.Left.Left = new(Node)
	head.Left.Right = new(Node)
	head.Right.Left = new(Node)
	head.Right.Right = new(Node)
	fmt.Println(IsFBT2(head))
}

func TestIsBalanced(t *testing.T) {
	head := new(Node)
	head.Left = new(Node)
	head.Right = new(Node)
	head.Left.Left = new(Node)
	head.Left.Right = new(Node)
	head.Right.Left = new(Node)
	head.Right.Right = new(Node)
	fmt.Println(IsBalanced(head))
}

func TestLowestCommonAncestor(t *testing.T) {
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
	head.Left.Left.Left = new(Node)
	head.Left.Left.Left.Data = 8
	head.Left.Left.Right = new(Node)
	head.Left.Left.Right.Data = 9
	head.Left.Right.Left = new(Node)
	head.Left.Right.Left.Data = 10
	head.Left.Right.Right = new(Node)
	head.Left.Right.Right.Data = 11
	fmt.Println(LowestCommonAncestor1(head, head.Left.Left, head.Left.Right).Data)           // expected output: 2
	fmt.Println(LowestCommonAncestor1(head, head.Left.Left.Left, head.Left.Right.Left).Data) // expected output: 2
	fmt.Println(LowestCommonAncestor1(head, head.Left.Left.Left, head.Left.Left).Data)       //expected output: 4

	fmt.Println(LowestCommonAncestor2(head, head.Left.Left, head.Left.Right).Data)           // expected output: 2
	fmt.Println(LowestCommonAncestor2(head, head.Left.Left.Left, head.Left.Right.Left).Data) // expected output: 2
	fmt.Println(LowestCommonAncestor2(head, head.Left.Left.Left, head.Left.Left).Data)       //expected output: 4
}

func TestGetSuccessorNode(t *testing.T) {
	head := new(NewNode)
	head.Data = 1
	head.Left = new(NewNode)
	head.Left.Data = 2
	head.Left.Parent = head
	head.Right = new(NewNode)
	head.Right.Data = 3
	head.Right.Parent = head
	head.Left.Left = new(NewNode)
	head.Left.Left.Data = 4
	head.Left.Left.Parent = head.Left
	head.Left.Right = new(NewNode)
	head.Left.Right.Data = 5
	head.Left.Right.Parent = head.Left
	head.Left.Left.Left = new(NewNode)
	head.Left.Left.Left.Data = 6
	head.Left.Left.Left.Parent = head.Left.Left
	head.Left.Left.Left.Left = new(NewNode)
	head.Left.Left.Left.Left.Data = 7
	head.Left.Left.Left.Left.Parent = head.Left.Left.Left
	head.Left.Left.Left.Right = new(NewNode)
	head.Left.Left.Left.Right.Data = 8
	head.Left.Left.Left.Right.Parent = head.Left.Left.Left
	/*
							1
				       2         3
		            4     5
				6
			7       8
	*/
	fmt.Println(GetSuccessorNode(head.Left.Left.Left.Right).Data) // expected output: 4
	fmt.Println(GetSuccessorNode(head.Left.Left.Left.Left).Data)  // expected output: 6
}

func TestMarshalBTToString(t *testing.T) {
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
	fmt.Println(MarshalBTToString(head))
}

func TestUnMarshalStringToBT(t *testing.T) {
	PreOrderTraversal1(UnMarshalStringToBT("1!2!4!*!*!5!*!*!3!6!*!*!7!*!*!"))
}

func TestPaperFolds(t *testing.T) {
	PaperFolds(2)
}
