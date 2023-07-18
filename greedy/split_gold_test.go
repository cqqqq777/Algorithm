package greedy

import (
	"fmt"
	"testing"
)

func TestSplitGold(t *testing.T) {
	fmt.Println(SplitGold([]int{15, 15, 20, 25, 30})) // expected output: 255
	fmt.Println(SplitGold([]int{10, 20, 30}))         // expected output: 90
}
