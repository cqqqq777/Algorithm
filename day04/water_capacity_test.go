package day04

import (
	"fmt"
	"testing"
)

func TestWaterCapacity(t *testing.T) {
	fmt.Println(WaterCapacity([]int{3, 1, 2, 5, 2, 4})) // expected output: 5
	fmt.Println(WaterCapacity([]int{4, 5, 1, 3, 2}))    // expected output: 2
}
