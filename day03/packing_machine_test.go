package day03

import (
	"fmt"
	"testing"
)

func TestPackingMachine(t *testing.T) {
	fmt.Println(PackingMachine([]int{1, 0, 5}))      // expected output: 3
	fmt.Println(PackingMachine([]int{4, 2, 0}))      // expected output: 2
	fmt.Println(PackingMachine([]int{100, 0, 0, 0})) // expected output: 75
}
