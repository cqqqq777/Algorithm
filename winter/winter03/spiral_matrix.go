package winter03

// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素
func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	r1, c1, r2, c2 := 0, 0, len(matrix)-1, len(matrix[0])-1
	res, i := make([]int, len(matrix)*len(matrix[0])), 0
	for r1 <= r2 && c1 <= c2 {
		row1, col1, row2, col2 := r1, c1, r2, c2
		if row1 == row2 {
			for j := col1; j <= col2; j++ {
				res[i] = matrix[row1][j]
				i++
			}
		} else if col1 == col2 {
			for j := row1; j <= row2; j++ {
				res[i] = matrix[j][col1]
				i++
			}
		} else {
			for j := col1; j < col2; j++ {
				res[i] = matrix[row1][j]
				i++
			}
			for j := row1; j < row2; j++ {
				res[i] = matrix[j][col2]
				i++
			}
			for j := col2; j > col1; j-- {
				res[i] = matrix[row2][j]
				i++
			}
			for j := row2; j > row1; j-- {
				res[i] = matrix[j][col1]
				i++
			}
		}
		r1++
		c1++
		r2--
		c2--
	}
	return res
}
