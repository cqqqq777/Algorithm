package winter22

// 给定一个三角形 triangle ，找出自顶向下的最小路径和。
// 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
// 也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	// 把三角形看成二叉树。
	// 就是求二叉树的最小路径和
	var recur func(idx, level int) int
	recur = func(idx, level int) int {
		if level == len(triangle)-1 {
			return triangle[level][idx]
		}

		return min(recur(idx, level+1), recur(idx+1, level+1)) + triangle[level][idx]
	}

	return recur(0, 0)
}

// 普通递归会超时
// 改为动态规划
// 发现只要 idx 和 level 确定，最小路径和就确定
// 第 i 层依赖与第 i+1 层
func dpVersion(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	// dp[level][idx] 表示位于第 level 层，索引为 idx 的点的最小路径和
	// level 从 0 开始，最大值为 len(triangle)
	dp := make([][]int, len(triangle))
	dp[len(triangle)-1] = make([]int, len(triangle[len(triangle)-1]))

	for idx := 0; idx < len(triangle[len(triangle)-1]); idx++ {
		dp[len(triangle)-1][idx] = triangle[len(triangle)-1][idx]
	}
	for level := len(triangle) - 2; level >= 0; level-- {
		dp[level] = make([]int, len(triangle[level]))
		for idx := 0; idx < len(triangle[level]); idx++ {
			dp[level][idx] = min(dp[level+1][idx], dp[level+1][idx+1]) + triangle[level][idx]
		}
	}
	return dp[0][0]
}
