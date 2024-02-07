package winter21

// 给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。
// 对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。
func longestIncreasingPath(matrix [][]int) int {
	var res int
	cache := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		cache[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			res = max(res, memorySearchV2(matrix, i, j, cache))
		}
	}
	return res
}

// 递归版本，会超时
func dfs(m [][]int, x, y, last int) int {
	if x < 0 || x == len(m) || y < 0 || y >= len(m[0]) || m[x][y] <= last {
		return 0
	}

	cur := m[x][y]
	// 不记录怎么走到这个位置的
	// 往回走了的话，在 dfs 入口处会直接返回
	return max(dfs(m, x+1, y, cur),
		dfs(m, x-1, y, cur),
		dfs(m, x, y+1, cur),
		dfs(m, x, y-1, cur),
	) + 1
}

// 记忆化搜索版本v1,会超内存，尝试优化 dfs 取消 last 参数
// cache[last][x][y] 表示上一个位置的值为 last，当前位置为 x, y 的最大递增路径
func memorySearch(m [][]int, x, y, last int, cache map[int]map[int][]int) int {
	if x < 0 || x == len(m) || y < 0 || y >= len(m[0]) || m[x][y] <= last {
		return 0
	}

	c, ok := cache[last]
	if !ok {
		c = make(map[int][]int)
		cache[last] = c
	}

	// 上一个位置的缓存
	up, ok := c[x-1]
	if !ok {
		up = make([]int, len(m[0])+1)
		c[x-1] = up
		// 特殊符号记录一下，表示这个位置还没有查找过
		c[x-1][y] = -1
	}

	if up[y] == -1 {
		up[y] = memorySearch(m, x-1, y, last, cache)
	}

	// 下一个位置的缓存
	down, ok := c[x+1]
	if !ok {
		down = make([]int, len(m[0])+1)
		c[x+1] = down
		c[x+1][y] = -1
	}

	if down[y] == -1 {
		down[y] = memorySearch(m, x+1, y, last, cache)
	}

	// 左一个位置和右一个位置的缓存
	row, ok := c[x]
	if !ok {
		row = make([]int, len(m[0])+1)
		c[x] = row
		row[y+1], row[y-1] = -1, -1
	}

	if row[y+1] == -1 {
		row[y+1] = memorySearch(m, x, y+1, last, cache)
	}
	if row[y-1] == -1 {
		row[y-1] = memorySearch(m, x, y-1, last, cache)
	}

	return max(up[y], down[y], row[y+1], row[y-1]) + 1
}

func memorySearchV2(m [][]int, x, y int, cache [][]int) int {
	if cache[x][y] != 0 {
		return cache[x][y]
	}

	var up, down, left, right int

	// 下一行的记录
	if x > 0 && m[x][y] > m[x-1][y] {
		down = memorySearchV2(m, x-1, y, cache)
	}

	// 上一行的记录
	if x < len(m)-1 && m[x][y] > m[x+1][y] {
		up = memorySearchV2(m, x+1, y, cache)
	}

	// 左边的记录
	if y > 0 && m[x][y] > m[x][y-1] {
		left = memorySearchV2(m, x, y-1, cache)
	}

	//右边的记录
	if y < len(m[0])-1 && m[x][y] > m[x][y+1] {
		right = memorySearchV2(m, x, y+1, cache)
	}

	res := max(left, right, up, down) + 1
	cache[x][y] = res
	return res
}
