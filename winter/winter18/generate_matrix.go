package winter18

// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n)
	}

	// 和螺旋打印类似
	row1, col1, row2, col2 := 0, 0, n-1, n-1
	val := 1
	for row1 <= row2 && col1 <= col2 {
		helper(matrix, row1, col1, row2, col2, &val)
		row1++
		col1++
		row2--
		col2--
	}
	return matrix
}

func helper(m [][]int, row1, col1, row2, col2 int, val *int) {
	if row1 == row2 {
		for col1 <= col2 {
			m[row1][col1] = *val
			col1++
			*val++
		}
	} else if col1 == col2 {
		for row1 <= row2 {
			m[row1][col1] = *val
			row1++
			*val++
		}
	} else {
		row, col := row1, col1
		for col < col2 {
			m[row1][col] = *val
			col++
			*val++
		}
		for row < row2 {
			m[row][col] = *val
			row++
			*val++
		}
		for col > col1 {
			m[row][col] = *val
			col--
			*val++
		}
		for row > row1 {
			m[row][col] = *val
			row--
			*val++
		}
	}
}
