package sort

import (
	"algorithm/heap"
	"math"
	"math/rand"
	"time"
)

func BubbleSort(sli []int) {
	if sli == nil || len(sli) < 2 {
		return
	}

	for i := 0; i < len(sli)-1; i++ {
		for j := 0; j < len(sli)-i-1; j++ {
			if sli[j] > sli[j+1] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
			}
		}
	}
}

func InsertionSort(sli []int) {
	if sli == nil || len(sli) < 2 {
		return
	}

	for i := 0; i < len(sli)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if sli[j] > sli[j-1] {
				break
			}
			sli[j], sli[j-1] = sli[j-1], sli[j]
		}
	}
}

func SelectionSort(sli []int) {
	if sli == nil || len(sli) < 2 {
		return
	}

	for i := 0; i < len(sli)-1; i++ {
		min := i
		for j := i + 1; j < len(sli); j++ {
			if sli[min] > sli[j] {
				sli[min], sli[j] = sli[j], sli[min]
			}
		}
	}
}

// QuickSort 3.0 版本，每次随机选取一个数作为基准
func QuickSort(sli []int) {
	quickSort(sli, 0, len(sli)-1)
}

func quickSort(arr []int, l, r int) {
	if l < r {
		rand.Seed(time.Now().UnixNano())
		randIndex := rand.Intn(r-l+1) + l
		arr[randIndex], arr[r] = arr[r], arr[randIndex]
		p := partition(arr, l, r)
		quickSort(arr, l, p[0]-1)
		quickSort(arr, p[1]+1, r)
	}
}

// 返回小于区的右边界和大于区的左边界
func partition(arr []int, l, r int) []int {
	less, more := l-1, r
	for l < more {
		if arr[l] < arr[r] {
			arr[l], arr[less+1] = arr[less+1], arr[l]
			less++
			l++
		} else if arr[l] > arr[r] {
			arr[l], arr[more-1] = arr[more-1], arr[l]
			more--
		} else {
			l++
		}
	}
	arr[r], arr[more] = arr[more], arr[r]
	return []int{less + 1, more}
}

func HeapSort(sli []int) {
	if sli == nil || len(sli) < 2 {
		return
	}

	for heapSize := heap.ConvertArrToHeap(sli, 0, len(sli)-1, heap.MaxHeapComparator()); heapSize > 1; {
		sli[0], sli[heapSize-1] = sli[heapSize-1], sli[0]
		heapSize--
		heap.Heapify(sli, 0, heapSize, heap.MaxHeapComparator())
	}
}

func ShellSort(sli []int) {}

func MergeSort(sli []int) {
	if sli == nil || len(sli) < 2 {
		return
	}

	process(sli, 0, len(sli)-1)
}

func process(arr []int, l, r int) {
	if l == r {
		return
	}
	mid := l + (r-l)>>1
	process(arr, l, mid)
	process(arr, mid+1, r)
	merge(arr, l, r, mid)
}

func merge(arr []int, l, r, mid int) {
	help := make([]int, r-l+1)
	i, p1, p2 := 0, l, mid+1
	for p1 <= mid && p2 <= r {
		if arr[p1] < arr[p2] {
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
		i++
	}
	for p1 <= mid {
		help[i] = arr[p1]
		p1++
		i++
	}
	for p2 <= r {
		help[i] = arr[p2]
		p2++
		i++
	}
	i = 0
	for l <= r {
		arr[l] = help[i]
		i++
		l++
	}
}

func RadixSort(sli []int) {
	if sli == nil || len(sli) < 2 {
		return
	}

	digit := maxBits(sli)
	bucket := make([]int, len(sli))

	for d := 1; d <= digit; d++ {
		count := make([]int, 10)
		for _, v := range sli {
			count[getDigit(v, d)]++
		}

		for i := 1; i < 10; i++ {
			count[i] += count[i-1]
		}

		for i := len(sli) - 1; i >= 0; i-- {
			j := getDigit(sli[i], d)
			bucket[count[j]-1] = sli[i]
			count[j]--
		}
		for i, _ := range sli {
			sli[i] = bucket[i]
		}
	}
}

func getDigit(x, d int) (digit int) {
	return (x / int(math.Pow10(d-1))) % 10
}

func maxBits(arr []int) (result int) {
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	for max != 0 {
		result++
		max /= 10
	}
	return
}

func BucketSort(sli []int) {}

func CountingSort(sli, r []int) {
	if sli == nil || len(sli) < 2 || len(r) < 2 {
		return
	}

	left, right := r[0], r[1]
	count := make([]int, right-left+1)

	for i := 0; i < len(sli); i++ {
		count[sli[i]-left]++
	}

	index := 0
	for i := 0; i < len(count); i++ {
		for j := count[i]; j > 0; j-- {
			sli[index] = i + left
			index++
		}
	}
}

func PdqSort(sli []int) {}
