package union_find_set

// Islands 判断有多少个岛
func Islands(matrix [][]int) int {
	if matrix == nil || matrix[0] == nil {
		return 0
	}
	width := len(matrix)
	length := len(matrix[0])
	num := 0
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			if matrix[i][j] == 1 {
				infect(matrix, i, j, width, length)
				num++
			}
		}
	}
	return num
}

func infect(matrix [][]int, i, j, width, length int) {
	if i < 0 || i >= width || j < 0 || j >= length || matrix[i][j] != 1 {
		return
	}
	matrix[i][j] = 2
	infect(matrix, i+1, j, width, length)
	infect(matrix, i-1, j, width, length)
	infect(matrix, i, j+1, width, length)
	infect(matrix, i, j-1, width, length)
}
