package day04

import "Algorithm/str"

// IsRotation 判断两个字符串是否互为旋转词
// 判断 b 是否为 a+a 子串即可（KMP）
func IsRotation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	return str.KMP(a+a, b) != -1
}
