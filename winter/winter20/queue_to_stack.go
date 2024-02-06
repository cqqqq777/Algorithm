package winter20

// 请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。
// 实现 MyStack 类：
// void push(int x) 将元素 x 压入栈顶。
// int pop() 移除并返回栈顶元素。
// int top() 返回栈顶元素。
// boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。

type Queue []int

func (q *Queue) Add(x int) {
	*q = append(*q, x)
}

func (q *Queue) Top() int {
	return (*q)[0]
}

func (q *Queue) Pop() int {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

type MyStack struct {
	q Queue
}

func Constructor() MyStack {
	return MyStack{
		q: make(Queue, 0),
	}
}

func (s *MyStack) Push(x int) {
	// 只要每次进队的时候，都进到队头就可以实现栈
	s.q.Add(x)
	for i := 0; i < len(s.q)-1; i++ {
		s.q.Add(s.q.Pop())
	}
}

func (s *MyStack) Pop() int {
	return s.q.Pop()
}

func (s *MyStack) Top() int {
	return s.q.Top()
}

func (s *MyStack) Empty() bool {
	return len(s.q) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
