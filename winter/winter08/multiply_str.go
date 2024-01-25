package winter08

import (
	"math"
	"strconv"
)

// 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
// 注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。
func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	if m < n {
		return multiply(num2, num1)
	}
	var res uint64
	// 长度短的每一位去乘长度长的数
	for i := n - 1; i >= 0; i-- {
		var carry, curRes uint64
		for j := m - 1; j >= 0; j-- {
			cur := uint64((num2[i]-'0')*(num1[j]-'0')) + carry
			carry = cur / 10
			curRes += (cur % 10) * uint64(math.Pow10(m-j-1))
		}
		if carry != 0 {
			curRes += carry * uint64(math.Pow10(m))
		}
		res += curRes * uint64(math.Pow10(n-i-1))
	}
	return strconv.FormatUint(res, 10)
}
