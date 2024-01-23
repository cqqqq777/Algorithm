package winter06

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
func generateParenthesis(n int) []string {
	res := make([]string, 0)

	var back func(subStr string, idx, flag int)
	back = func(subStr string, idx, flag int) {
		if flag < 0 {
			return
		}
		if idx == 2*n {
			if flag == 0 {
				res = append(res, subStr)
			}
			return
		}

		// 选左括号
		flag++
		subStr += "("
		back(subStr, idx+1, flag)
		flag--
		subStr = subStr[:len(subStr)-1]

		flag--
		subStr += ")"
		back(subStr, idx+1, flag)
		flag++
		subStr = subStr[:len(subStr)-1]
	}

	var track string
	back(track, 0, 0)
	return res
}
