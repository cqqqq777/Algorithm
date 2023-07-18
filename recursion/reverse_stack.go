package recursion

type Stack struct{}

func (s *Stack) Push(_ int) {}

func (s *Stack) Pop() int { return 0 }

func (s *Stack) IsEmpty() bool { return true }

// ReverseStack 逆序栈
func ReverseStack(stack *Stack) {
	if stack == nil || stack.IsEmpty() {
		return
	}
	last := GetAndRemoveLast(stack)
	ReverseStack(stack)
	stack.Push(last)
}

func GetAndRemoveLast(stack *Stack) int {
	cur := stack.Pop()
	if !stack.IsEmpty() {
		last := GetAndRemoveLast(stack)
		stack.Push(cur)
		return last
	}
	return cur
}
