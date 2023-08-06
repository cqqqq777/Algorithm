package day04

import (
	"fmt"
	"testing"
)

func TestGetMinStack(t *testing.T) {
	s := NewGetMinStack()
	s.Push(3)
	s.Push(23)
	s.Push(66)
	fmt.Println(s.GetMin()) // expected output: 3
	s.Push(1)
	fmt.Println(s.GetMin()) // expected output: 1
	fmt.Println(s.Pop())    // expected output: 1
	fmt.Println(s.GetMin()) // expected output: 3
	fmt.Println(s.Pop())    // expected output: 66
}
