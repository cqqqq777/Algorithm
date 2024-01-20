package day03

// TopK 给定字符串数组，返回出现次数最多的前 k 个
// 先做 map 做词频统计
// 大根堆：按大根堆组织后弹出 k 个
// 小根堆：维护一个数量为 k 的小根堆，只有当前词频大于堆顶的元素时，弹出堆顶，当前元素入堆
func TopK(arr []string, k int) []WordCount {
	wr := make(map[string]int)
	for _, v := range arr {
		wr[v] += 1
	}
	h := newHeap()
	for s, v := range wr {
		if h.size < k {
			h.heapInsert(WordCount{s: s, count: v})
		} else if h.Peek().count < v {
			h.Pop()
			h.heapInsert(WordCount{s: s, count: v})
		}
	}
	res := make([]WordCount, k)
	for i := 0; i < k; i++ {
		res[i] = h.Pop()
	}
	return res
}

type WordCount struct {
	s     string
	count int
}

// 为此题定制的一个小根堆
type heap struct {
	arr  []WordCount
	size int
}

func newHeap() *heap {
	return &heap{arr: make([]WordCount, 0)}
}

func (h *heap) heapify(index int) {
	if index < 0 || index >= h.size {
		return
	}
	left := index<<1 + 1
	for left < h.size {
		tmp := 0
		if left+1 < h.size && h.arr[left+1].count > h.arr[left].count {
			tmp = left + 1
		} else {
			tmp = left
		}
		if h.arr[tmp].count >= h.arr[index].count {
			break
		}
		h.arr[index], h.arr[tmp] = h.arr[tmp], h.arr[index]
		index = tmp
		left = index<<1 + 1
	}
}

func (h *heap) heapInsert(w WordCount) {
	h.size++
	h.arr = append(h.arr, w)
	index := h.size - 1
	for h.arr[index].count < h.arr[(index-1)/2].count {
		h.arr[index], h.arr[(index-1)/2] = h.arr[(index-1)/2], h.arr[index]
		index = (index - 1) / 2
	}
}

func (h *heap) Pop() WordCount {
	res := h.arr[0]
	h.arr[0], h.arr[h.size-1] = h.arr[h.size-1], h.arr[0]
	h.arr = h.arr[:h.size-1]
	h.size--
	h.heapify(0)
	return res
}

func (h *heap) Peek() *WordCount {
	if h.size > 0 {
		return &h.arr[0]
	}
	return nil
}
