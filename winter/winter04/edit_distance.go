package winter04

// 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数。
// 你可以对一个单词进行如下三种操作：
// 插入一个字符
// 删除一个字符
// 替换一个字符
func minDistance(word1, word2 string) int {
	if word1 == "" {
		return len(word2)
	}
	if word2 == "" {
		return len(word1)
	}

	// dp[i][j] 表示 word1 以 i-1 位置结尾，word2 以 j-1 位置结尾，所需要的最少编辑距离
	// word 的 -1 位置表示 ""
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word2); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for i := 0; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if i == 0 {
				dp[i][j] = j
				continue
			}
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
