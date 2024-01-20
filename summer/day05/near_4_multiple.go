package day05

// Near4Multiple 给定一个数组判断其是否能调整为任意两个相邻的数相乘是4的倍数
// 统计 arr 中奇数的个数和偶数的个数，偶数又分为两类统计，一种为4倍数的数，一种不为4倍数的数
func Near4Multiple(arr []int) bool {
	if arr == nil || len(arr) < 1 {
		return false
	}
	// a, b, c 分别统计三种书
	a, b, c := 0, 0, 0
	for i := 0; i < len(arr); i++ {
		if arr[i]&1 == 1 {
			a++
		} else if a%4 == 0 {
			c++
		} else {
			b++
		}
	}
	if b == 0 {
		if a == 1 {
			return c >= 1
		} else {
			return a >= c-1
		}
	} else if b == 1 {
		return c >= 1 && c >= a
	} else {
		return c >= a
	}
}
