package day05

// DeleteWood
// 在迷迷糊糊的大草原上，小红捡到了n根木棍，第i根木棍的长度为i，小红现在很开心。想选出其中的三根木棍组成美丽的三角形。但是小明想捉弄小红，想去掉一些木棍，使得小红任意选三根木棍都不能组成三角形。
// 请问小明最少去掉多少根木棍呢？
// 给定N，返回至少去掉多少根？
// 最多保留 n 范围内不算第一项的情况下斐波那契数的个数
func DeleteWood(n int) int {
	var res int
	for res = 1; ; res++ {
		if fib(res) > n {
			break
		}
	}
	return n - res + 2
}
