package day04

import (
	"fmt"
	"testing"
)

func TestStackConvertQueue(t *testing.T) {
	s := NewStackConvertQueue()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	for i := 0; i < 6; i++ {
		fmt.Println(s.Pop())
	}
}

func TestQueueConvertStack(t *testing.T) {

}
