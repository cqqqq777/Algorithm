package str

import (
	"math"
	"strings"
)

// Manacher 获取最大回文子串
func Manacher(str string) int {
	if len(str) < 1 {
		return 0
	}
	if len(str) == 1 {
		return 1
	}
	str = handleStr(str)            // 123 -> #1#2#1#
	R, C, max := -1, -1, 0          // R 为最右边界+1，max 为经处理后的 str 最大回文半径
	lenArr := make([]int, len(str)) // 记录每一位字符的回文半径
	for i := 0; i < len(str); i++ {
		if R > i {
			lenArr[i] = int(math.Min(float64(2*C-i), float64(R-i)))
		} else {
			lenArr[i] = 1
		}
		for i+lenArr[i] < len(str) && i-lenArr[i] > -1 {
			if str[i+lenArr[i]] == str[i-lenArr[i]] {
				lenArr[i] += 1
			} else {
				break
			}
		}
		if i+lenArr[i] > R {
			R = i + lenArr[i]
			C = i
		}
		max = int(math.Max(float64(max), float64(lenArr[i])))
	}
	return max - 1
}

func handleStr(str string) string {
	var builder strings.Builder
	builder.Grow(2*len(str) + 1)
	spe := '#'
	for _, v := range str {
		builder.WriteRune(spe)
		builder.WriteRune(v)
	}
	builder.WriteRune(spe)
	return builder.String()
}
