package day04

import "math"

// MaxAbsBetweenLeftRight 给定数组长度 N
// 将任意长度大于0小于 N 的前缀作为左部分，剩下的为右部分，返回左部分最大值减右部分最大值的最大绝对值
func MaxAbsBetweenLeftRight(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	Max := math.MinInt
	for i := 0; i < len(arr); i++ {
		if arr[i] > Max {
			Max = arr[i]
		}
	}
	return Max - min(arr[0], arr[len(arr)-1])
}
