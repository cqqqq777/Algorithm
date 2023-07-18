package list

import (
	"fmt"
	"testing"
)

func TestPalindromeList2(t *testing.T) {
	head := new(Node)
	head.Data = 1
	head.Next = new(Node)
	head.Next.Data = 2
	head.Next.Next = new(Node)
	head.Next.Next.Data = 3
	head.Next.Next.Next = new(Node)
	head.Next.Next.Next.Data = 2
	head.Next.Next.Next.Next = new(Node)
	head.Next.Next.Next.Next.Data = 1
	fmt.Println(PalindromeList2(head))
}
