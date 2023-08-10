package day07

// FindMissingNum
// 给定一个整数数组A，长度为n，有 1 <= A[i] <= n，且对于[1,n]的整数，其中部分整数会重复出现而部分不会出现。
// 实现算法找到[1,n]中所有未出现在A中的整数。
// 提示：尝试实现O(n)的时间复杂度和O(1)的空间复杂度（返回值不计入空间复杂度）。
// 输入描述：
// 一行数字，全部为整数，空格分隔
// A0 A1 A2 A3... 输出描述：
// 一行数字，全部为整数，空格分隔R0 R1 R2 R3... 示例1:
// 输入
// 1 3 4 3
// 输出
// 2
// 尝试将 i 位置都放上 i+1 的数，最后遍历一遍数组如果 i 位置的值不为 i+1，那 i+1 就是缺失的数
func FindMissingNum(arr []int) []int {
	if arr == nil || len(arr) == 0 {
		return []int{}
	}
	res := make([]int, 0)
	for _, v := range arr {
		modify(arr, v)
	}
	for k, v := range arr {
		if v != k+1 {
			res = append(res, k+1)
		}
	}
	return res
}

func modify(arr []int, value int) {
	for arr[value-1] != value {
		t := arr[value-1]
		arr[value-1] = value
		value = t
	}
}
