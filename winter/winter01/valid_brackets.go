package winter01

type Stack []int32

func (s *Stack) Push(ele int32) {
	*s = append(*s, ele)
}

func (s *Stack) Pop() (int32, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r, true
}

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
func isValid(s string) bool {
	if s == "" || len(s)&1 == 1 {
		return false
	}

	stack := make(Stack, 0)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack.Push(v)
			continue
		}
		if r, ok := stack.Pop(); !ok || r != m[v] {
			return false
		}
	}
	return len(stack) == 0
}

var m = map[int32]int32{')': '(', '}': '{', ']': '['}
