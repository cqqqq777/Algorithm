package day03

// FindNumInSortedMatrix 判断 n 是否存在于有序 matrix 中
func FindNumInSortedMatrix(matrix [][]int, n int) bool {
	row, col := 0, len(matrix[0])-1
	for row < len(matrix) && col >= 0 {
		if n == matrix[row][col] {
			return true
		} else if n < matrix[row][col] {
			col--
		} else {
			row++
		}
	}
	return false
}
