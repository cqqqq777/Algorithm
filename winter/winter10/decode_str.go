package winter10

import (
	"strconv"
	"strings"
)

// 给定一个经过编码的字符串，返回它解码后的字符串。
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

type Stack []int

func (s *Stack) Push(element int) {
	*s = append(*s, element)
}

func (s *Stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func decodeString(s string) string {
	if len(s) <= 1 {
		return s
	}
	stack := make(Stack, 0)
	flag := true
	var builder strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if flag {
				stack.Push(i)
				flag = false
			}
			continue
		}
		if s[i] == '[' {
			stack.Push(i)
			flag = true
			continue
		}
		// p1, p2, i 分别指向数字开头, '[', ']'
		if s[i] == ']' {
			p2, p1 := stack.Pop(), stack.Pop()
			num, _ := strconv.Atoi(s[p1:p2])
			builder.Grow(num * (i - p2 - 1))
			for j := 0; j < num; j++ {
				builder.WriteString(s[p2+1 : i])
			}
			cur := builder.String()
			builder.Reset()
			s = s[0:p1] + cur + s[i+1:]
			// i 注意要加上切片前后的长度差值
			i += (num * (i - p2 - 1)) - (i - p1 + 1) - 1
		}
	}
	return s
}
