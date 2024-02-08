package winter22

// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
// 大根堆：弹 k 个
// 小根堆：维护大小为 k 的小根堆，只有当新元素的词频大于堆顶时，弹出堆顶，新元素入堆
func topKFrequent(nums []int, k int) []int {
	// 词频统计
	countMap := make(map[int]int)
	for _, v := range nums {
		countMap[v]++
	}
	heap := make([]int, 0, k)
	flag := 0
	for key, v := range countMap {
		// 先加 k 个元素到答案数组里面去
		if flag < k {
			heap = append(heap, key)
			flag++
			continue
		}
		// 建堆
		if flag == k {
			convertToHeap(heap, countMap)
			flag++
		}

		if v > countMap[heap[0]] {
			heap[0] = key
			heapify(heap, countMap)
		}
	}

	return heap
}

func convertToHeap(nums []int, countMap map[int]int) {
	for size := 1; size < len(nums); size++ {
		heapInsert(nums, size, countMap)
	}
}

// heapify 是自顶向下的调整过程
// 此处将堆顶元素向下调整
func heapify(nums []int, countMap map[int]int) {
	idx := 0
	for idx < len(nums) {
		left := idx<<1 + 1
		if left >= len(nums) {
			break
		}
		right := left + 1
		if right >= len(nums) {
			right = left
		}
		minIdx := idx
		if countMap[nums[left]] < countMap[nums[minIdx]] {
			minIdx = left
		}
		if countMap[nums[right]] < countMap[nums[minIdx]] {
			minIdx = right
		}
		if minIdx == idx {
			break
		}
		nums[idx], nums[minIdx] = nums[minIdx], nums[idx]
		idx = minIdx
	}
}

func heapInsert(nums []int, idx int, countMap map[int]int) {
	for idx > 0 {
		father := (idx - 1) >> 1
		if countMap[nums[father]] > countMap[nums[idx]] {
			nums[father], nums[idx] = nums[idx], nums[father]
			idx = father
		} else {
			break
		}
	}
}
