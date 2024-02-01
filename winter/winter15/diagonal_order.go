package winter15

// 给你一个大小为 m x n 的矩阵 mat ，请以对角线遍历的顺序，用一个数组返回这个矩阵中的所有元素。
func findDiagonalOrder(mat [][]int) []int {
	row1, col1, row2, col2 := 0, 0, 0, 0
	flag := false
	res := make([]int, 0)
	// 每次都打印对角线
	// 用两个点分别表示对角线的端点
	for row1 < len(mat) {
		diagonalPrint(mat, row1, col1, row2, col2, flag, &res)
		if col1 == len(mat[0])-1 {
			row1++
		} else {
			col1++
		}
		if row2 == len(mat)-1 {
			col2++
		} else {
			row2++
		}
		flag = !flag
	}
	return res
}

func diagonalPrint(m [][]int, row1, col1, row2, col2 int, flag bool, res *[]int) {
	// flag 为 true 时，从上往下打印，否则从下往上打印
	if flag {
		for col1 >= col2 {
			*res = append(*res, m[row1][col1])
			row1++
			col1--
		}
	} else {
		for col1 >= col2 {
			*res = append(*res, m[row2][col2])
			row2--
			col2++
		}
	}
}
