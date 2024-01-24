package winter07

// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
// 你可以认为每种硬币的数量是无限的。
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// dp[i] 表示换 i 所需要的最少硬币数
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]] != -1 {
				if dp[i] == -1 {
					dp[i] = dp[i-coins[j]] + 1
				} else {
					dp[i] = min(dp[i], dp[i-coins[j]]+1)
				}
			}
		}
	}
	return dp[amount]
}
