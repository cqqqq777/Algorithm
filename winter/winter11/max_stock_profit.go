package winter11

// 给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
// 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
// 返回 你能获得的 最大 利润 。
func maxProfit(prices []int) int {
	// dp[i][0] 表示第 i 天没有股票的最大利润
	// dp[i][1] 表示第 i 天有股票的最大利润
	dp := make([][2]int, len(prices))
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][0]
}
