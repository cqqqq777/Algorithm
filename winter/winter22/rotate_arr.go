package winter22

// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
func rotate(nums []int, k int) {
	if len(nums) == 1 {
		return
	}

	k = k % len(nums)
	if k == 0 {
		return
	}

	// 先全部翻转
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}

	// 再分别翻转前面的前 k 个和后 len(nums)-k 个
	for i := 0; i < k/2; i++ {
		nums[i], nums[k-1-i] = nums[k-1-i], nums[i]
	}

	for i := 0; i < (len(nums)-k)/2; i++ {
		nums[k+i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[k+i]
	}
}
