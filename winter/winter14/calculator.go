package winter14

// 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
// 注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func calculate(s string) int {
	return helper(&s)
}

func helper(s *string) int {
	var sign uint8 = '+'
	num := 0
	stack := make(Stack, 0)
	for len(*s) > 0 {
		c := (*s)[0]
		*s = (*s)[1:]
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		}

		if c == '(' {
			num = helper(s)
		}

		if ((c < '0' || c > '9') && c != ' ') || len(*s) == 0 {
			if sign == '+' {
				stack.Push(num)
			} else if sign == '-' {
				stack.Push(-num)
			}
			sign = c
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
