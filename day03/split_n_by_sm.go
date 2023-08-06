package day03

// SplitNBySM 定义两种操作，求最少操作数到达 n
// 定义 s == m = 1
// 操作1：m = s   s *= 2
// 操作2：s = m + s
// 对于质数只执行操作二就是最优解
// 任意合数都可以写成质数相乘的形式
func SplitNBySM(n int) int {
	if n < 2 {
		return 0
	}
	if isPrime(n) {
		return n - 1
	}
	sumAndCount := divSumAndCount(n)
	return sumAndCount[0] - sumAndCount[1]
}

// 返回一个合数分解为质数相乘时，所有因子的和与因子的个数
func divSumAndCount(n int) []int {
	sum, count := 0, 0
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			sum += i
			count++
			n /= i
		}
	}
	return []int{sum, count}
}

func isPrime(n int) bool {
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
