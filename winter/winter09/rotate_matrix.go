package winter09

// 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
func rotate(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 || len(matrix) != len(matrix[0]) {
		return
	}
	row1, col1, row2, col2 := 0, 0, len(matrix)-1, len(matrix)-1
	for row1 < row2 && col1 < col2 {
		rotate1(matrix, row1, col1, row2, col2)
		row1++
		col1++
		row2--
		col2--
	}
}

func rotate1(m [][]int, row1, col1, row2, col2 int) {
	for i := 0; i < col2-col1; i++ {
		m[row1][col1+i], m[row1+i][col2] = m[row1+i][col2], m[row1][col1+i]
		m[row1][col1+i], m[row2][col2-i] = m[row2][col2-i], m[row1][col1+i]
		m[row1][col1+i], m[row2-i][col1] = m[row2-i][col1], m[row1][col1+i]
	}
}
