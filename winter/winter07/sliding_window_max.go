package winter07

// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
// 你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	dequeue := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		// 维护单调递减队列
		for len(dequeue) != 0 && nums[dequeue[len(dequeue)-1]] <= nums[i] {
			dequeue = dequeue[:len(dequeue)-1]
		}
		dequeue = append(dequeue, i)
		// 代表此时的队首已经不在窗口内了
		if dequeue[0] == i-k {
			dequeue = dequeue[1:]
		}
		// 表示已经达到窗口大小可以开始弹出元素了
		if i >= k-1 {
			res = append(res, nums[dequeue[0]])
		}
	}
	return res
}
