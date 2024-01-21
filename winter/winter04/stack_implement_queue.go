package winter04

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

type MyQueue struct {
	s1, s2 Stack
}

func Constructor() MyQueue {
	return MyQueue{
		s1: make(Stack, 0),
		s2: make(Stack, 0),
	}
}

func (q *MyQueue) Push(x int) {
	q.s1.Push(x)
}

func (q *MyQueue) Pop() int {
	if len(q.s2) == 0 {
		for len(q.s1) > 0 {
			q.s2.Push(q.s1.Pop())
		}
	}
	return q.s2.Pop()
}

func (q *MyQueue) Peek() int {
	if len(q.s2) == 0 {
		for len(q.s1) > 0 {
			q.s2.Push(q.s1.Pop())
		}
	}
	return q.s2[len(q.s2)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.s1) == 0 && len(q.s2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
