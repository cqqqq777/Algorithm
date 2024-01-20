package dp

// MinDistance 给定两个字符串，返回 word1 变为 word2 的最少操作数，每次仅可操作一个字符
func MinDistance(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}
	dp := make([][]int, len(word1)+1)
	dp[0] = make([]int, len(word2)+1)
	for i := 0; i < len(word2)+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
		for j := 0; j < len(word2)+1; j++ {
			if j == 0 {
				dp[i][j] = i
				continue
			}
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i][j-1], dp[i-1][j-1]), dp[i-1][j]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
