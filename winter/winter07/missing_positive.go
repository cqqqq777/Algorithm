package winter07

// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = n + 1
		}
	}

	for i := 0; i < n; i++ {
		// 有可能前面的数会把后面的数改了
		// 到后面的时候要改回来，故用绝对值
		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
