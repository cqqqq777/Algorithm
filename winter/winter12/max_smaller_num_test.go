package winter12

import (
	"fmt"
	"testing"
)

func TestMaxSmaller(t *testing.T) {
	fmt.Println(maxSmaller([]int{1, 2, 4, 9}, 2499))      // 2494
	fmt.Println(maxSmaller([]int{1, 2, 4, 9}, 4999))      // 4994
	fmt.Println(maxSmaller([]int{1, 2, 4, 9}, 9999))      // 9994
	fmt.Println(maxSmaller([]int{1, 2, 4, 9}, 836778683)) // 499999999
	fmt.Println(maxSmaller([]int{1}, 10))                 // 1
	fmt.Println(maxSmaller([]int{9}, 99999))              // 9999
}
