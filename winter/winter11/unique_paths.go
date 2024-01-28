package winter11

// 一个机器人位于一个 m * n 网格的左上角 （起始点在下图中标记为 “Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
// 问总共有多少条不同的路径？
func uniquePaths(m int, n int) int {
	if m == n && m == 1 {
		return 1
	}
	// dp[i][j] 表示以矩阵 i，j 位置为起点，有多少条路
	dp := make([][]int, m)
	dp[m-1] = make([]int, n)
	for i := n - 2; i >= 0; i-- {
		dp[m-1][i] = 1
	}
	for i := m - 2; i >= 0; i-- {
		dp[i] = make([]int, n)
		dp[i][n-1] = 1
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = dp[i+1][j] + dp[i][j+1]
		}
	}
	return dp[0][0]
}
