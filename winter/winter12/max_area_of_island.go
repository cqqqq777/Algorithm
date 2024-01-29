package winter12

// 给你一个大小为 m x n 的二进制矩阵 grid 。
// 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
// 岛屿的面积是岛上值为 1 的单元格的数目。
// 计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
func maxAreaOfIsland(grid [][]int) int {
	var res int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				var area int
				infect(grid, i, j, &area)
				res = max(res, area)
			}
		}
	}
	return res
}

func infect(grid [][]int, x, y int, area *int) {
	if x < 0 || x == len(grid) || y < 0 || y == len(grid[0]) || grid[x][y] != 1 {
		return
	}
	grid[x][y] = 2
	*area++
	// 分别感染上下左右
	infect(grid, x-1, y, area)
	infect(grid, x+1, y, area)
	infect(grid, x, y-1, area)
	infect(grid, x, y+1, area)
}
