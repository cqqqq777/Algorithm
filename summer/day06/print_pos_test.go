package day06

import (
	"fmt"
	"testing"
)

func TestPrintPos(t *testing.T) {
	fmt.Println(PrintPos([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 2, 5, 1, 6, 3, 7})) // expected output: 4 5 2 6 7 3 1
}
