package day01

// MaxBorderSize 给定矩阵只有0和1，返回边框全是1的最大正方形的边
func MaxBorderSize(matrix [][]int) int {
	return maxBorderSize1(matrix)
}

// 暴力枚举，前两重循环确定点的位置，第三重循环确定边长,第四重循环判断这一点以此为边是否能构成正方形
// 边长的范围由先到边界的方向确定
func maxBorderSize1(matrix [][]int) int {
	N, M := len(matrix), len(matrix[0])
	result := 0
	for row := 0; row < N; row++ {
		for col := 0; col < M; col++ {
		Lab:
			for size := 1; size <= min(N-row, M-col); size++ {
				// 验证上下两边是否满足条件
				for start := col; start < col+size; start++ {
					if matrix[row][start] != 1 || matrix[row+size-1][start] != 1 {
						continue Lab
					}
				}
				// 验证左右两边
				for start := row; start < row+size; start++ {
					if matrix[start][col] != 1 || matrix[start][col+size-1] != 1 {
						continue Lab
					}
				}
				if result < size {
					result = size
				}
			}
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 先对矩阵做预处理，可以在确定点和边长的情况下快速判断是否构成正方形
// 定义 right[i][j] 记录右边连续出现1的个数，down[i][j] 记录下方连续出现1的个数
func maxBorderSize2(matrix [][]int) int {
	right, down := preHandleMatrix(matrix)
	M, N := len(matrix), len(matrix[0])
	result := 0
	for row := 0; row < M; row++ {
		for col := 0; col < N; col++ {
			for size := min(M-row, N-col); size >= 1; size-- {
				if right[row][col] >= size && down[row][col] >= size && right[row+size-1][col] >= size && down[row][col+size-1] >= size {
					if result < size {
						result = size
					}
					break
				}
			}
		}
	}
	return result
}

func preHandleMatrix(matrix [][]int) ([][]int, [][]int) {
	M, N := len(matrix), len(matrix[0])
	right, down := make([][]int, M), make([][]int, M)
	right[M-1] = make([]int, N)
	down[M-1] = make([]int, N)
	if matrix[M-1][N-1] == 1 {
		right[M-1][N-1] = 1
		down[M-1][N-1] = 1
	}
	// 处理最后一列
	for i := M - 2; i >= 0; i-- {
		right[i] = make([]int, N)
		down[i] = make([]int, N)
		if matrix[i][N-1] == 1 {
			right[i][N-1] = 1
			down[i][N-1] = down[i+1][N-1] + 1
		}
	}
	// 处理最后一行
	for i := N - 2; i >= 0; i-- {
		if matrix[M-1][i] == 1 {
			right[M-1][i] = right[M-1][i+1] + 1
			down[M-1][i] = 1
		}
	}
	for row := M - 2; row >= 0; row-- {
		for col := N - 2; col >= 0; col-- {
			if matrix[row][col] == 1 {
				right[row][col] = right[row][col+1] + 1
				down[row][col] = down[row+1][col] + 1
			}
		}
	}
	return right, down
}
