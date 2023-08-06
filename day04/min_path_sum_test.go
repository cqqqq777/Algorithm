package day04

import (
	"fmt"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	fmt.Println(MinPathSum([][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 1, 1, 1, 1},
	}))
}
