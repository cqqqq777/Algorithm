package winter12

// 给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
// 测试用例的答案是一个 32-位 整数。
// 子数组 是数组的连续子序列。
func maxProduct(nums []int) int {
	maxF, minF, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], mn*nums[i], nums[i])
		minF = min(mx*nums[i], mn*nums[i], nums[i])
		res = max(maxF, res)
	}
	return res
}
