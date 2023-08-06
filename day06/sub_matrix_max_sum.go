package day06

import "math"

// SubMatrixMaxSum
// 给定一个整型矩阵，返回子矩阵的最大累计和。
func SubMatrixMaxSum(matrix [][]int) int {
	res := math.MinInt
	for i := 0; i < len(matrix); i++ {
		help := make([]int, len(matrix[0]))
		for j := i; j < len(matrix); j++ {
			for k := 0; k < len(matrix[0]); k++ {
				help[k] += matrix[j][k]
			}
			m := RecentMaxScore(help)
			if m > res {
				res = m
			}
		}
	}
	return res
}
