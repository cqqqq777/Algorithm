package winter18

// 给你一个满足下述两条属性的 m x n 整数矩阵：
// 每行中的整数从左到右按非严格递增顺序排列。
// 每行的第一个整数大于前一行的最后一个整数。
// 给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。
func searchMatrix(matrix [][]int, target int) bool {
	// 看右上角，可以看作二叉搜索树
	row, col := 0, len(matrix[0])-1
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] < target {
			row++
		} else if matrix[row][col] > target {
			col--
		} else {
			return true
		}
	}
	return false
}
