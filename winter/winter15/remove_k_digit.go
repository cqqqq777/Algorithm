package winter15

import (
	"strings"
)

// 给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。
// 请你以字符串形式返回这个最小的数字。
func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}

	// 维护单调栈，保证左边的数字尽可能小
	res := make([]byte, 0, len(num)-k)
	for i := range num {
		v := num[i]
		for k > 0 && len(res) > 0 && res[len(res)-1] > v {
			res = res[:len(res)-1]
			k--
		}
		res = append(res, v)
	}
	res = res[:len(res)-k]
	n := strings.TrimLeft(string(res), "0")
	if n == "" {
		return "0"
	}
	return n
}
