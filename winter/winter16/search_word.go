package winter16

// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
// 同一个单元格内的字母不允许被重复使用。
func exist(board [][]byte, word string) bool {
	// 和岛屿问题类似
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if infect(board, word, i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

func infect(m [][]byte, word string, i, j, idx int) bool {
	if idx == len(word) {
		return true
	}
	if i < 0 || i == len(m) || j < 0 || j == len(m[0]) || m[i][j] == '0' {
		return false
	}
	if m[i][j] != word[idx] {
		return false
	}
	element := m[i][j]
	m[i][j] = '0'
	flag := infect(m, word, i+1, j, idx+1) || infect(m, word, i-1, j, idx+1) || infect(m, word, i, j+1, idx+1) || infect(m, word, i, j-1, idx+1)
	m[i][j] = element
	return flag
}
