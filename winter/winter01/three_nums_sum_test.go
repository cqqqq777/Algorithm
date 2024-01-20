package winter01

import (
	"fmt"
	"testing"
)

func TestThreeNums(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
