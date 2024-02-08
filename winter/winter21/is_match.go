package winter21

// 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
// 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
func isMatch(s string, p string) bool {
	// 状态：
	// 大小为len(s)+1 * len(p)+1的状态数组dp
	// dp[i][j]是bool类型，表示s的前i个字母组成的子字符串s[0,i-1]，是否与p的前j个字母组成的子字符串p[0,j-1]匹配
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}

	// 起始状态：
	// 当i==0且j==0时，此时两个串都是空串，此时一定匹配，即dp[0][0]=true
	dp[0][0] = true
	// 当i!=0且j==0时，表示字符串p是空串，由于字符串s非空，所以一定匹配不上，所以dp[i][0]=false
	// go的slice中默认值为false，此处显式复制只是为了更直观，可以省略；后序赋值false的情况同理
	for i := 1; i < m+1; i++ {
		dp[i][0] = false
	}
	// 当i==0且j!=0时，表示字符串s是空串
	for j := 1; j < n+1; j++ {
		if p[j-1] == '*' {
			// 如果p[j-1]==*，由于*可以使得前面的字符个数为0，所以p[j-2,j-1]可以与空串匹配，此时只需要判断p[0,j-3]是否与空串s匹配即可，因此dp[0][j]=dp[0][j-2]
			dp[0][j] = dp[0][j-2]
		} else {
			// 如果p[j-1]!=*，则字符串p[0,j-1]一定非空，空串s无法与非空的p匹配,dp[0][j]=false
			dp[0][j] = false
		}
	}

	// 3.状态转移
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				// 如果最右端字符匹配（s[i-1]==p[j-1] 或 p[j-1]==.）：
				// 由于最右端的字符是匹配的，此时将s[i-1]和p[j-1]抵消，只需要看前面的其他字符是否匹配，即看s[0,i-2]与p[0,j-2]是否匹配，因此dp[i][j]=dp[i-1][j-1]
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 如果最右端字符不匹配
				if p[j-1] == '*' {
					// 如果p[j-1]==*，则可能匹配
					if s[i-1] == p[j-2] || p[j-2] == '.' {
						// 如果字符s[i-1]和p[j-2]匹配（s[i-1]==p[j-2] 或 p[j-2]==.） ，此时要考虑以下3种情况：
						// p[j-1]的通配符*将p[j-2]字符重复了0次，因此p[j-2,j-1]自动消除，只需要比较s[0,i-1]和p[0,j-3]是否匹配即可，因此dp[i][j]=dp[i][j-2]
						// p[j-1]的通配符*将p[j-2]字符重复了1次，因此s[i-1]与p[j-2,j-1]相互抵消，只需要比较s[0,i-2]和p[0,j-3]是否匹配即可，因此dp[i][j]=dp[i-1][j-2]
						// p[j-1]的通配符*将p[j-2]字符重复了2次或2次以上，此时将p[j-2,j-1]产生的多个重复字母中的其中一个重复字母与s[i-1]相互抵消，只需要比较s[0,i-2]和p[0,j-1]是否匹配即可，因此dp[i][j]=dp[i-1][j]
						dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j]
					} else {
						// 如果字符s[i-1]和p[j-2]不匹配：
						// 则使用p[j-1]的*将p[j-2]重复0次，使得p[j-2,j-1]自动消除，只需要比较s[0,i-1]和p[0,j-3]是否匹配即可，因此dp[i][j]=dp[i][j-2]
						dp[i][j] = dp[i][j-2]
					}
				} else {
					// 如果p[j-1]!=*，则两个字符串一定不匹配，dp[i][j]=false
					dp[i][j] = false
				}
			}
		}
	}

	// 当规模递增到i=len(s)，j=len(p)时的子问题的解就是全局问题的解
	return dp[m][n]
}