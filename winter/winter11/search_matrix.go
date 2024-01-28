package winter11

// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
func searchMatrix(matrix [][]int, target int) bool {
	row, col := 0, len(matrix[0])-1
	// 对于每一个数，左边都比自身小，下边都比自身大
	// 当作二叉搜索树去找 target
	for row < len(matrix) && col >= 0 {
		if target == matrix[row][col] {
			return true
		} else if target < matrix[row][col] {
			col--
		} else {
			row++
		}
	}
	return false
}
