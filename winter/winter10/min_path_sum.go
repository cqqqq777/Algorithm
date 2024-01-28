package winter10

// 给定一个包含非负整数的 m * n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// dp[i][j] 表示以 grid[i][j] 为起点的最小路径
	// dp[i][j] = min(dp[i+1][j], dp[i][j+1]
	dp := make([][]int, m)
	dp[m-1] = make([]int, n)
	dp[m-1][n-1] = grid[m-1][n-1]
	for i := m - 1; i >= 0; i-- {
		if i != m-1 {
			dp[i] = make([]int, n)
		}
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				continue
			}
			if i+1 < m && j+1 < n {
				dp[i][j] = min(dp[i+1][j], dp[i][j+1]) + grid[i][j]
			} else if i+1 < m {
				dp[i][j] = dp[i+1][j] + grid[i][j]
			} else {
				dp[i][j] = dp[i][j+1] + grid[i][j]
			}
		}
	}
	return dp[0][0]
}
