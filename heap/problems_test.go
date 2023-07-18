package heap

import (
	"fmt"
	"testing"
)

func TestSortedArrDistanceLessK(t *testing.T) {
	sli := []int{3, 2, 1, 6, 5, 4, 9, 8, 7, 12, 11, 10}
	SortedArrDistanceLessK(sli, 2)
	fmt.Println(sli)
}
