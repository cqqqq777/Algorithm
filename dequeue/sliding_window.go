package dequeue

// GetMaxWindow 获取每个滑动窗口最大值
func GetMaxWindow(arr []int, w int) []int {
	if w > len(arr) || w < 1 || arr == nil {
		return []int{}
	}
	res := make([]int, 0, len(arr)-w+1)
	dequeue := NewDequeue()
	for i := 0; i < len(arr); i++ {
		for !dequeue.IsEmpty() && arr[dequeue.PeekAtTail()] <= arr[i] {
			dequeue.PoolAtTail()
		}
		dequeue.AddAtTail(i)
		if dequeue.PeekAtHead() == i-w {
			dequeue.PoolAtHead()
		}
		if i >= w-1 {
			res = append(res, arr[dequeue.PeekAtHead()])
		}
	}
	return res
}
