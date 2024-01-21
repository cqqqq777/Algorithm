package winter04

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
func lengthOfLIS(nums []int) int {
	if nums == nil {
		return 0
	}

	//dp[i] 表示必须以 i 位置结尾，最长递增子序列的长度
	dp := make([]int, len(nums))
	dp[0] = 1
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		m := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				m = max(dp[j], m)
			}
		}
		dp[i] = m + 1
		res = max(dp[i], res)
	}
	return res
}

// 最长递增子序列的二分版本，时间复杂度 O(N * logN)
func lengthOfLIS2(nums []int) int {
	if nums == nil {
		return 0
	}

	// 使用 maxV[i] 表示长度为 i 的递增子序列的最大元素的最小值
	// 遍历 nums 时，二分法找到 nums[i] 在 maxV 中小于其的最大值，记录其下标，看能否让 maxV[i+1] 更新
	// 最后 maxV 的长度就是最长递增子序列的长度
	maxV := make([]int, len(nums))
	maxV[0] = nums[0]
	idx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxV[idx] {
			idx++
			maxV[idx] = nums[i]
		} else {
			pos := bs(maxV, idx, nums[i])
			maxV[pos] = min(nums[i], maxV[pos])
		}
	}
	return idx + 1
}

func bs(maxV []int, r, target int) int {
	left, right := 0, r
	for left <= right {
		mid := (right-left)>>1 + left
		if maxV[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right + 1
}
