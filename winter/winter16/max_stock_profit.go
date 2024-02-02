package winter16

// 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
func maxProfit(prices []int) int {
	// dp[i][k][0] 表示第 i 天，手上没有股票，且交易的限制数为 k 的情况下的最大利润
	// dp[i][k][1] 表示第 i 天，手上持有股票，且交易的限制数为 k 的情况下的最大利润
	dp := make([][3][2]int, len(prices))
	dp[0][1][1] = -prices[0]
	dp[0][2][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		// 前一天有可能持有股票，有可能没有股票
		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1]+prices[i])
		dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1]+prices[i])

		dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0]-prices[i])
		dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0]-prices[i])
	}
	return dp[len(prices)-1][2][0]
}
