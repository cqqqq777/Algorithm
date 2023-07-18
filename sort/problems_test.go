package sort

import (
	"fmt"
	"testing"
)

func TestMinSum(t *testing.T) {
	fmt.Println(MinSum([]int{1, 3, 4, 2, 5}))
}

func TestReverseOrderPair(t *testing.T) {
	fmt.Println(ReverseOrderPair([]int{1, 2, 1, 2, 1}))
}

func TestPartition1(t *testing.T) {
	sli := []int{4, 1, 2, 3, 2, 8, 5}
	Partition1(sli, 5)
	fmt.Println(sli)
}

func TestPartition2(t *testing.T) {
	sli := []int{4, 1, 2, 3, 2, 8, 5}
	Partition2(sli, 2)
	fmt.Println(sli)
}
