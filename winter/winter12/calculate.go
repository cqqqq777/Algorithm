package winter12

// 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
// 整数除法仅保留整数部分。
// 你可以假设给定的表达式总是有效的。所有中间结果将在 [-231, 231 - 1] 的范围内。
// 注意：不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

// 注意如果有括号需要递归的情况，要传入 s 的指针
func calculate(s string) int {
	stack := make(Stack, 0)
	sign := '+'
	num := 0

	for len(s) > 0 {
		c := s[0]
		s = s[1:]
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		}

		if c == '(' {
			num = calculate(s)
		}

		if ((c < '0' || c > '9') && c != ' ') || len(s) == 0 {
			if sign == '+' {
				stack.Push(num)
			} else if sign == '-' {
				stack.Push(-num)
			} else if sign == '*' {
				val := stack.Pop()
				stack.Push(val * num)
			} else if sign == '/' {
				val := stack.Pop()
				stack.Push(val / num)
			}
			sign = int32(c)
			num = 0
		}
		if c == ')' {
			break
		}
	}

	var res int
	for _, v := range stack {
		res += v
	}
	return res
}
