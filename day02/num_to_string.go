package day02

import (
	"strconv"
)

// NumToString 给定数字返回将其转换为字符串的方法数
func NumToString(num int) int {
	if num <= 0 {
		return 0
	}
	str := strconv.Itoa(num)
	return f(str, 0)
}

func f(s string, index int) int {
	if index == len(s) {
		return 1
	}
	if s[index] == '0' {
		return 0
	}
	res := f(s, index+1)
	if index == len(s)-1 {
		return res
	}
	if (s[index]-'0')*10+(s[index+1]-'0') < 27 {
		res += f(s, index+2)
	}
	return res
}

func NumToString2(num int) int {
	if num <= 0 {
		return 0
	}
	str := strconv.Itoa(num)
	return fDP(str)
}

// dp 版本
func fDP(s string) int {
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1
	if s[len(s)-1] != '0' {
		dp[len(s)-1] = 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] != '0' {
			dp[i] = dp[i+1]
			if (s[i]-'0')*10+(s[i+1]-'0') < 27 {
				dp[i] += dp[i+2]
			}
		}
	}
	return dp[0]
}
