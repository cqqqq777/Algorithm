package winter06

// 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
// 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
// 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	l, r := 0, x
	res := -1
	for l <= r {
		mid := (r-l)>>1 + l
		if mid*mid <= x {
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}
