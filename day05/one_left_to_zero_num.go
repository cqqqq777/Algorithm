package day05

// OneLeftToZero
// 字符串只由'0'和'1'两种字符构成，
// 当字符串长度为1时，所有可能的字符串为"0"、"1"；当字符串长度为2时，所有可能的字符串为"00"、"01"、"10"、"11"；当字符串长度为3时，所有可能的字符串为"000"、"001"、"010"、"011"、"100"、"101"、"110"、"111"
// 如果某一个字符串中，只要是出现'0'的位置，左边就靠着'1'，这样的字符串叫作达标字符串。
// 给定一个正数N，返回所有长度为N的字符串中，达标字符串的数量。比如，N=3，返回3，因为只有"101"、"110"、"111"达标。
func OneLeftToZero(n int) int {
	return oneLeftToZeroProcess(n)
}

func oneLeftToZeroProcess(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	return oneLeftToZeroProcess(n-1) + oneLeftToZeroProcess(n-2)
}

// 只要递归严格依赖以前的项，都可以优化为O(log N)
// 将斐波那契数列优化到O(log N)
func fib(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	base := [][]int{
		{1, 1},
		{1, 0},
	}
	base = matrixPow(base, n-2)
	return base[0][0] + base[1][0]
}

// 将求矩阵的 n 次方优化到O(log N)
func matrixPow(matrix [][]int, n int) [][]int {
	res := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		res[i] = make([]int, len(matrix[0]))
		res[i][i] = 1
	}
	tmp := matrix
	for ; n != 0; n >>= 1 {
		if n&1 != 0 {
			res = multiMatrix(res, tmp)
		}
		tmp = multiMatrix(tmp, tmp)
	}
	return res
}

// return m1 * m2
func multiMatrix(m1, m2 [][]int) [][]int {
	res := make([][]int, len(m1))
	for i := 0; i < len(m2[0]); i++ {
		res[i] = make([]int, len(m2[0]))
	}
	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m2[0]); j++ {
			for k := 0; k < len(m2); k++ {
				res[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return res
}
