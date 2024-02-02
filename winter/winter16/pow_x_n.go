package winter16

// 实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	var helper func(num float64, p int) float64
	helper = func(num float64, p int) float64 {
		if p == 0 {
			return 1
		}
		y := helper(num, p/2)
		if p%2 == 0 {
			return y * y
		}
		return y * y * num
	}
	if n > 0 {
		return helper(x, n)
	} else {
		return 1.0 / helper(x, -n)
	}
}
