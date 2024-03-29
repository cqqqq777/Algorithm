package winter02

// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
// 输入：[7,1,5,3,6,4]
// 输出：5
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
// 注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
func maxProfit(prices []int) int {
	if prices == nil {
		return 0
	}
	// dp[i] 表示必须在第 i 天买入股票，最大的利润
	// dp[i] + prices[i] 就表示 i 以后的最大价格
	// 故 dp[i-1] = dp[i] + prices[i] - prices[i-1] (如果小于 0 就直接等于 0)
	dp := make([]int, len(prices))
	var res int
	for i := len(prices) - 2; i >= 0; i-- {
		dp[i] = max(dp[i+1]+prices[i+1]-prices[i], 0)
		res = max(dp[i], res)
	}
	return res
}
