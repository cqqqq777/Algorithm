package day02

// ParenthesisDepth 给定合法括号串计算深度
func ParenthesisDepth(s string) int {
	if s == "" {
		return 0
	}
	res := 0
	count := 0
	for _, v := range s {
		if v == '(' {
			count++
		} else {
			count--
		}
		if count > res {
			res++
		}
	}
	return res
}

// MaxLength 给定只包含括号的字符串，找出最长合法括号字符串长度
func MaxLength(s string) int {
	if s == "" {
		return 0
	}
	dp := make([]int, len(s))
	pre, res := 0, 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			pre = i - dp[i-1] - 1
			if pre >= 0 && s[pre] == '(' {
				dp[i] += dp[i-1] + 2
				if pre > 0 {
					dp[i] += dp[pre-1]
				}
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
