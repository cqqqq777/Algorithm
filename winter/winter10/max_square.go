package winter10

// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
func maximalSquare(matrix [][]byte) int {
	right, down := preHandle(matrix)
	var res int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != '1' {
				continue
			}
			edge := 1
			minDown := down[i][j]
			for k := 1; k < right[i][j]; k++ {
				minDown = min(minDown, down[i][j+k])
				if minDown < k+1 {
					break
				}
				edge++
			}
			res = max(edge, res)
		}
	}
	return res * res
}

func preHandle(matrix [][]byte) ([][]int, [][]int) {
	// right, down 分别记录 matrix[i][j] 右边和下边有多少个连续的 1（包含自身）
	right, down := make([][]int, len(matrix)), make([][]int, len(matrix))

	m, n := len(matrix), len(matrix[0])
	right[m-1] = make([]int, n)
	down[m-1] = make([]int, n)

	if matrix[m-1][n-1] == '1' {
		right[m-1][n-1] = 1
		down[m-1][n-1] = 1
	}

	// 处理最后一行
	for col := n - 2; col >= 0; col-- {
		if matrix[m-1][col] == '1' {
			right[m-1][col] = right[m-1][col+1] + 1
			down[m-1][col] = 1
		}
	}

	// 处理最后一列
	for row := m - 2; row >= 0; row-- {
		right[row] = make([]int, n)
		down[row] = make([]int, n)
		if matrix[row][n-1] == '1' {
			down[row][n-1] = down[row+1][n-1] + 1
			right[row][n-1] = 1
		}
	}

	for row := m - 2; row >= 0; row-- {
		for col := n - 2; col >= 0; col-- {
			if matrix[row][col] == '1' {
				right[row][col] = right[row][col+1] + 1
				down[row][col] = down[row+1][col] + 1
			}
		}
	}

	return right, down
}

func dpVersion(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	var res int

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 如果当前位置的值为 '1'
			if matrix[i-1][j-1] == '1' {
				// 计算当前位置的 dp 值，取左上、上、左三个相邻位置的最小值加一
				dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
				res = max(res, dp[i][j])
			}
		}
	}

	// 最大正方形的面积为边长的平方
	return res * res
}
