package day06

import (
	"fmt"
	"testing"
)

func TestSubMatrixMaxSum(t *testing.T) {
	fmt.Println(SubMatrixMaxSum([][]int{
		{-5, 3, 6, 4},
		{-7, 9, -5, 3},
		{-7, 1, -200, 4},
	})) // expected output: 20
}
