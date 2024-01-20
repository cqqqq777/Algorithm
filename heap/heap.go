package heap

type Comparator func(int, int) bool

type Comparators[T comparable] func(a, b T) bool

func MaxHeapComparator() Comparator {
	return func(a int, b int) bool {
		return a > b
	}
}

func MinHeapComparator() Comparator {
	return func(a int, b int) bool {
		return a < b
	}
}

func InsertHeap(arr []int, index int, comparator Comparator) {
	for comparator(arr[index], arr[(index-1)/2]) {
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
}

func Heapify(arr []int, index, size int, comparator Comparator) {
	left := index<<1 + 1
	for left < size {
		// tmp used to record the bigger or smaller index
		tmp := 0
		if left+1 < size && !comparator(arr[left], arr[left+1]) {
			tmp = left + 1
		} else {
			tmp = left
		}
		if comparator(arr[index], arr[tmp]) {
			break
		}
		arr[index], arr[tmp] = arr[tmp], arr[index]
		index = tmp
		left = index<<1 + 1
	}
}

func ConvertArrToHeap(arr []int, left, right int, comparator Comparator) int {
	index, heapSize := left, left
	for index <= right {
		InsertHeap(arr, index, comparator)
		heapSize++
		index = heapSize
	}
	return right - left + 1
}
