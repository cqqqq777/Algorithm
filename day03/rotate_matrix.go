package day03

// RotateMatrix 将矩阵每个数顺时针旋转90°
func RotateMatrix(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 || len(matrix) != len(matrix[0]) {
		return
	}
	row1, col1, row2, col2 := 0, 0, len(matrix)-1, len(matrix)-1
	for row1 < row2 && col1 < col2 {
		rotate(matrix, row1, col1, row2, col2)
		row1++
		col1++
		row2--
		col2--
	}
}

func rotate(m [][]int, row1, col1, row2, col2 int) {
	for i := 0; i < col2-col1; i++ {
		m[row1][col1+i], m[row1+i][col2] = m[row1+i][col2], m[row1][col1+i]
		m[row1][col1+i], m[row2][col2-i] = m[row2][col2-i], m[row1][col1+i]
		m[row1][col1+i], m[row2-i][col1] = m[row2-i][col1], m[row1][col1+i]
	}
}
