package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func generateSli() []int {
	rand.Seed(time.Now().UnixNano())
	sli := make([]int, 100)
	for i := 0; i < 100; i++ {
		sli = append(sli, rand.Intn(1000))
	}
	return sli
}

func TestBubbleSort(t *testing.T) {
	sli := generateSli()
	BubbleSort(sli)
	fmt.Println(sli)
}

func TestInsertionSort(t *testing.T) {
	sli := generateSli()
	InsertionSort(sli)
	fmt.Println(sli)
}

func TestSelectionSort(t *testing.T) {
	sli := generateSli()
	SelectionSort(sli)
	fmt.Println(sli)
}

func TestMergeSort(t *testing.T) {
	sli := generateSli()
	MergeSort(sli)
	fmt.Println(sli)
}

func TestHeapSort(t *testing.T) {
	sli := generateSli()
	HeapSort(sli)
	fmt.Println(sli)
}

func TestQuickSort(t *testing.T) {
	sli := generateSli()
	QuickSort(sli)
	fmt.Println(sli)
}

func TestCountingSort(t *testing.T) {
	sli := generateSli()
	CountingSort(sli, []int{0, 1000})
	fmt.Println(sli)
}

func TestRadixSort(t *testing.T) {
	sli := generateSli()
	RadixSort(sli)
	fmt.Println(sli)
}
