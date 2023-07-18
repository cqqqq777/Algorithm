package heap

// SortedArrDistanceLessK 给定一个几乎有序的数组排好序，几乎有序是指给定一个 k，每个元素距离它原本的位置位置不超过 k
func SortedArrDistanceLessK(arr []int, k int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	heapSize := 0
	if len(arr) <= k {
		heapSize = len(arr)
	} else {
		heapSize = k + 1
	}

	// create heap
	heap := make([]int, heapSize)
	i := 0
	for ; i < heapSize; i++ {
		heap[i] = arr[i]
	}
	ConvertArrToHeap(heap, 0, heapSize-1, MinHeapComparator())

	i = 0
	index := heapSize
	for index < len(arr) {
		arr[i] = heap[0]
		heap[0] = arr[index]
		Heapify(heap, 0, heapSize, MinHeapComparator())
		index++
		i++
	}

	for ; i < len(arr); i++ {
		arr[i] = heap[0]
		heapSize--
		heap[0], heap[heapSize] = heap[heapSize], heap[0]
		Heapify(heap, 0, heapSize, MinHeapComparator())
	}
}
