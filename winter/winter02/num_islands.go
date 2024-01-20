package winter02

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。
// 输入：grid = [
//
//	["1","1","1","1","0"],
//	["1","1","0","1","0"],
//	["1","1","0","0","0"],
//	["0","0","0","0","0"]
//
// ]
// 输出：1
// 每次遇到 1 就把周围的 1 全部感染成 2
func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				infect(grid, i, j)
				res++
			}
		}
	}
	return res
}

func infect(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i == len(grid) || j == len(grid[i]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = 2
	infect(grid, i, j-1)
	infect(grid, i, j+1)
	infect(grid, i+1, j)
	infect(grid, i-1, j)
}
