package winter21

// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
// 同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	// 0 和 n-1 的房间选一个打劫，那退化成普通打家劫舍
	var robCommon func(nums []int) int
	robCommon = func(prices []int) int {
		if len(prices) == 1 {
			return prices[0]
		}
		if len(prices) == 2 {
			return max(prices[0], prices[1])
		}
		dp := make([]int, len(prices))
		dp[0] = prices[0]
		dp[1] = max(prices[0], prices[1])
		for i := 2; i < len(prices); i++ {
			dp[i] = max(dp[i-2]+prices[i], dp[i-1])
		}
		return dp[len(prices)-1]
	}
	return max(robCommon(nums[1:]), robCommon(nums[:len(nums)-1]))
}
