package winter01

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 子数组 是数组中的一个连续部分。
func maxSubArray(nums []int) int {
	if nums == nil {
		return 0
	}
	// dp[i] 表示必须以 i 位置结尾，连续子数组的最大和
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = nums[i] + dp[i-1]
		}
		res = max(dp[i], res)
	}
	return res
}
