package dequeue

import (
	"fmt"
	"testing"
)

func TestGetMaxWindow(t *testing.T) {
	fmt.Println(GetMaxWindow([]int{4, 3, 5, 4, 3, 3, 6, 7, 8}, 3))
}
