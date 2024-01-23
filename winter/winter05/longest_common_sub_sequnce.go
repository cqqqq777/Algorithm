package winter05

// 最长公共子序列问题
func longestCommonSubsequence(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}

	// dp[i][j] 代表必须 text1 以 i-1 位置结尾，text2 以 j - 1 位置结尾，的最长公共子序列
	// 如果 text1[i] == text2[j]，那么 dp[i][j] 等于 dp[i-1][j-1] + 1
	// 否则等于 max(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]
	dp := make([][]int, len(text1)+1)
	dp[0] = make([]int, len(text2)+1)
	for i := 1; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}
