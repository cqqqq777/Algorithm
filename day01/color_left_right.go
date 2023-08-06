package day01

import "math"

// ColorLeftRight 给定一个数组只包含1，2，可以相互转换，要求最终1全在2前面，求最少的转换次数
func ColorLeftRight(arr []int) int {
	return colorLeftRight1(arr)
}

// 暴力枚举：左侧可以有 0..len(arr) 个1，计算每种可能的转换数最后返回最小值
// 假设在左侧有 i 个1，就是统计 i 以前有多少个2，i 以后有多少个1\
// 时间复杂度 O(N^2)
func colorLeftRight1(arr []int) int {
	min := math.MaxInt
	for i := 0; i <= len(arr); i++ {
		num := 0
		for j := 0; j < i; j++ {
			if arr[j] == 2 {
				num++
			}
		}
		for j := i; j < len(arr); j++ {
			if arr[j] == 1 {
				num++
			}
		}
		if num < min {
			min = num
		}
	}
	return min
}

// 在暴力枚举的内层循环中，每次都是统计1和2的数量，预处理数组使其不必循环即可得到结果
// 类似前缀和，时间复杂度 O(N)
func colorLeftRight2(arr []int) int {
	min := math.MaxInt
	// right 记录从某位置开始往后有多少个1
	right := make([]int, len(arr)+1)
	if arr[len(arr)-1] == 1 {
		right[len(arr)-1] = 1
	}
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] == 1 {
			right[i] = 1
		}
		right[i] += right[i+1]
	}
	// left 记录某位置以前有多少个2，因为与遍历顺序相同，用一个变量记录即可
	left := 0
	for i := 0; i <= len(arr); i++ {
		if i > 0 && i < len(arr) && arr[i] == 2 {
			left += 1
		}
		num := left + right[i]
		if min > num {
			min = num
		}
	}
	return min
}
