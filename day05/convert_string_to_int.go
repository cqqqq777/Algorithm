package day05

import (
	"math"
)

// ConvertStringToInt
// 定一个字符串，如果该字符串符合人们日常书写一个整数的形式，返回int类型的这个数；如果不符合或者越界返回-1或者报错。
func ConvertStringToInt(s string) (int, bool) {
	if s == "" {
		return 0, false
	}
	if !isValid(s) {
		return 0, false
	}
	var pois bool
	if s[0] == '-' {
		pois = true
	}
	i := 0
	if pois {
		i = 1
	}
	// 使用负数表示
	// 因为 int 的负数范围的绝对值比整数多一个
	res := 0
	for ; i < len(s); i++ {
		cur := -int(s[i] - '0')
		// 先判断是否越界再累加
		if res < math.MinInt/10 || (cur < (math.MinInt%10) && res == math.MinInt/10) {
			return 0, false
		}
		res = res*10 + cur
	}
	if !pois {
		if res == math.MinInt {
			return 0, false
		}
		return -res, true
	}
	return res, true
}

// 判断某个字符是否合法
func isValid(s string) bool {
	// 除了数字之外只有减号
	// 减号出现在开头且后接不为0的数字
	// 开头为0则后面不能有数字
	if s[0] != '-' && (s[0] < '0' || s[0] > '9') {
		return false
	}
	if s[0] == '-' && (s[1] == '0' || len(s) == 1) {
		return false
	}
	if s[0] == '0' && len(s) != 1 {
		return false
	}
	for i := 1; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}
