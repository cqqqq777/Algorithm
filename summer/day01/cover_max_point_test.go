package day01

import (
	"fmt"
	"testing"
)

func TestCoverMaxPoint(t *testing.T) {
	fmt.Println(CoverMaxPoint([]int{0, 13, 24, 35, 46, 57, 60, 72, 87}, 6))
	fmt.Println(CoverMaxPoint([]int{1, 2, 3, 4, 7, 8, 9, 10}, 4))
}
