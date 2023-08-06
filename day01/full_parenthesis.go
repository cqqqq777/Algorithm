package day01

// FullParenthesis 给定一个括号字符串，返回想要将其补充完整最少需要添加的括号数
func FullParenthesis(s string) int {
	if s == "" {
		return 0
	}
	count, res := 0, 0
	for _, v := range s {
		if v == '(' {
			count++
		}
		if v == ')' {
			count--
		}
		if count == -1 {
			count++
			res++
		}
	}
	res += count
	return res
}
