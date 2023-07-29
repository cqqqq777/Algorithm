package bitwise_operate

// GetMax 获取a，b中的较大值
func GetMax(a, b int) int {
	c := a - b
	sa := sign(c)
	sb := flip(sa)
	return sa*a + sb*b
}

// 负数返回 0
// 正数返回 1
func sign(n int) int {
	return flip(n >> 63 & 1)
}

// n 只能是1或0
// n == 1（0），返回0（1）
func flip(n int) int {
	return n ^ 1
}

// Is2Power 判断一个数是否为2的幂
func Is2Power(n int) bool {
	//return (n-1)&n == 0
	return n&(^n+1) == n
}

// Is4Power 判断一个数是否为4的幂
func Is4Power(n int) bool {
	return n&(^n+1) == n && n&(0x55555555<<1) == 0
}

// Plus return a + b
func Plus(a, b int) int {
	sum := a
	for b != 0 {
		sum = a ^ b
		b = (a & b) << 1
		a = sum
	}
	return sum
}

// Subtract return a-b
func Subtract(a, b int) int {
	return Plus(a, getOpposite(b))
}

func getOpposite(n int) int {
	return Plus(^n, 1)
}

// Multiply return a* b
func Multiply(a, b int) int {
	res := 0
	for b != 0 {
		if b&1 != 0 {
			res = Plus(res, a)
		}
		a <<= 1
		b >>= 1
	}
	return res
}

// Divide return a / b
func Divide(a, b int) int {
	if b == 0 {
		panic("divisor is 0")
	}
	res := 0
	for i := 63; i >= 0; i = Subtract(i, 1) {
		if a>>i >= b {
			res |= 1 << i
			a = Subtract(a, b<<i)
		}
	}
	if sign(a) == sign(b) {
		return res
	} else {
		return getOpposite(res)
	}
}
