package winter20

import "strings"

// 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
// 字母和数字都属于字母数字字符。
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	p1, p2 := 0, len(s)-1
	for p1 < p2 {
		var left byte
		if (s[p1] >= '0' && s[p1] <= '9') || (s[p1] >= 'a' && s[p1] <= 'z') {
			left = s[p1]
		} else {
			p1++
			continue
		}

		var right byte
		if (s[p2] >= '0' && s[p2] <= '9') || (s[p2] >= 'a' && s[p2] <= 'z') {
			right = s[p2]
		} else {
			p2--
			continue
		}

		if left != right {
			return false
		}
		p1++
		p2--
	}
	return true
}
