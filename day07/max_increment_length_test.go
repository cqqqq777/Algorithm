package day07

import (
	"fmt"
	"testing"
)

func TestMaxIncrementLength(t *testing.T) {
	fmt.Println(MaxIncrementLength([]int{3, 1, 2, 6, 3, 4}))                      // expected output: 4
	fmt.Println(MaxIncrementLength2([]int{3, 1, 2, 6, 3, 4}))                     // expected output: 4
	fmt.Println(MaxIncrementLength([]int{3, 1, 4, 6, 3, 2, 1, 3, 4, 6, 99, 56}))  // expected output: 6
	fmt.Println(MaxIncrementLength2([]int{3, 1, 4, 6, 3, 2, 1, 3, 4, 6, 99, 56})) // expected output: 6
	fmt.Println(MaxIncrementLength2([]int{1, 2, 3, 4, 5, 6}))                     // expected output: 6
}
