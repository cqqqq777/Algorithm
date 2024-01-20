package day01

import (
	"fmt"
	"testing"
)

func TestColorLeftRight(t *testing.T) {
	fmt.Println(ColorLeftRight([]int{2, 2, 1, 2, 2}))
	fmt.Println(colorLeftRight2([]int{2, 2, 1, 2, 2}))
}
