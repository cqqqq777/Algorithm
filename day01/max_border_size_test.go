package day01

import (
	"fmt"
	"testing"
)

func TestMaxBorderSize(t *testing.T) {
	fmt.Println(maxBorderSize2([][]int{
		{0, 1, 1, 1, 1, 1},
		{0, 1, 0, 0, 1, 1},
		{0, 1, 0, 0, 1, 1},
		{0, 1, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1},
	}))
}
