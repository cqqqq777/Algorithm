package day04

import "container/list"

// StackConvertQueue 使用栈实现队列
type StackConvertQueue struct {
	push, pop *list.List
}

func NewStackConvertQueue() *StackConvertQueue {
	return &StackConvertQueue{push: list.New(), pop: list.New()}
}

func (s *StackConvertQueue) Push(data int) {
	s.push.PushBack(data)
}

func (s *StackConvertQueue) Pop() (int, bool) {
	e := s.pop.Back()
	if e != nil {
		s.pop.Remove(e)
		return e.Value.(int), true
	}
	e = s.push.Back()
	for e != nil {
		s.pop.PushBack(e.Value.(int))
		s.push.Remove(e)
		e = s.push.Back()
	}
	e = s.pop.Back()
	if e == nil {
		return 0, false
	}
	s.pop.Remove(e)
	return e.Value.(int), true
}

// 仅使用队列实现栈类似
// 两个队列，push 时将所有元素 push 进一个队列，pop 时将有数据的队列转移到另一个队列剩一个数据返回
