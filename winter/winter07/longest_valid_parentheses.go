package winter07

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}

	// dp[i] 表示必须以 i 位置结尾的最长合法括号子串的长度
	dp := make([]int, len(s))

	if s[1] == ')' && s[0] == '(' {
		dp[1] = 2
	}
	res := dp[1]
	for i := 2; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] = dp[i-2] + 2
			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		res = max(res, dp[i])
	}
	return res
}
