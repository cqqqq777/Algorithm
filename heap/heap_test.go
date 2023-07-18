package heap

import (
	"fmt"
	"testing"
)

func TestHeapInsert(t *testing.T) {
	arr := []int{1, 2, 3, 2, 3}
	InsertHeap(arr, 4, MaxHeapComparator())
	fmt.Println(arr)
}

func TestHeapify(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	Heapify(arr, 0, 5, MaxHeapComparator())
	fmt.Println(arr)
}
