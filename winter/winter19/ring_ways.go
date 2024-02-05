package winter19

// 圆环上有10个点，编号为0~9。从0点出发，每次可以逆时针和顺时针走一步，问走n步回到0点共有多少种走法。
func ringWays(points, n int) int {
	if points == 0 || n == 0 {
		return 1
	}
	// dp[i][j] 表示走 i 步到 j 点的方案数
	dp := make([][]int, n+1)
	dp[0] = make([]int, points)
	dp[0][0] = 1
	for i := 1; i < len(dp); i++ {
		for j := 0; j < points; j++ {
			// 走 i 步到 j 等于走 i-1 步到 j 的相邻位置
			dp[i][j] = dp[i-1][(j-1+points)%points] + dp[i-1][(j+1+points)%points]
		}
	}
	return dp[n][0]
}
