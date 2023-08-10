package day07

import "math"

// MaxIncrementLength 最长递增子序列问题
// O(N^2)
func MaxIncrementLength(arr []int) int {
	if arr == nil {
		return 0
	}
	size := len(arr)
	// 以 i 位置结尾的最长递增子序列
	dp := make([]int, size)
	dp[0] = 1
	res := dp[0]
	for i := 1; i < size; i++ {
		dp[i] = 1
		m := math.MinInt
		for j := i - 1; j >= 0; j-- {
			if arr[j] < arr[i] && dp[j] > m {
				m = dp[j]
			}
		}
		if m != math.MinInt {
			dp[i] += m
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

// MaxIncrementLength2
// O(N * logN)
func MaxIncrementLength2(arr []int) int {
	// ends 数组分为有效区和无效区
	// 有效区的 i 位置表示所有长度为 i+1 的递增子序列中，最小的结尾的值
	// 在遍历 arr 时，每次去有效区查找大于 arr[i] 的最左位置，如果找到就更新，未找到就扩充有效区
	// 更新 dp 就是在 ends 数组中有效区内小于等于 arr[i] 数的个数
	size := len(arr)
	dp, ends := make([]int, size), make([]int, size)
	dp[0] = 1
	res := dp[0]
	// 有效区的结尾
	valid := 0
	ends[valid] = arr[0]
	for i := 1; i < size; i++ {
		// 二分查找大于 arr[i] 的最左位置
		l, r := 0, valid
		for l < r {
			mid := l + (r-l)>>1
			if ends[mid] >= arr[i] {
				r = mid
			} else {
				l = mid + 1
			}
		}
		if l == valid && ends[l] < arr[i] {
			valid++
			ends[valid] = arr[i]
			dp[i] = l + 2
		} else {
			ends[l] = arr[i]
			dp[i] = l + 1
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
