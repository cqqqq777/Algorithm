package winter06

// 整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。
// 更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。
// 如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	// 找到第一个升序
	for ; i >= 0 && nums[j] <= nums[i]; i, j = i-1, j-1 {
	}

	if i >= 0 {
		// 找到 i 第一个比 nums[i] 大的数
		for ; nums[i] >= nums[k]; k-- {
		}

		nums[i], nums[k] = nums[k], nums[i]
	}

	// 此时 i 后面是降序，将其变为升序，则字典序最小
	for i, j = j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
