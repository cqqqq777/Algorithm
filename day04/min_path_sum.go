package day04

// MinPathSum 给定 matrix 从左上走到右下，返回最小路径和
// 空间压缩的动态规划
func MinPathSum(matrix [][]int) int {
	dp := make([]int, len(matrix[0]))
	dp[0] = matrix[0][0]
	row, col := len(matrix), len(matrix[0])
	for i := 1; i < col; i++ {
		dp[i] = matrix[0][i] + matrix[0][i-1]
	}
	for i := 1; i < row; i++ {
		dp[0] = matrix[i][0] + dp[0]
		for j := 1; j < col; j++ {
			dp[j] = matrix[i][j] + min(dp[j-1], dp[j])
		}
	}
	return dp[col-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
