package day04

import "container/list"

// GetMinStack , Pop，Push，GetMin 操作时间复杂度为O(1)
// 维护两个栈，一个存放正常数据，一个存放 min
type GetMinStack struct {
	data, min *list.List
}

func NewGetMinStack() *GetMinStack {
	return &GetMinStack{data: list.New(), min: list.New()}
}

func (g *GetMinStack) Push(data int) {
	g.data.PushBack(data)
	e := g.min.Back()
	if e == nil {
		g.min.PushBack(data)
		return
	}
	m := e.Value.(int)
	if data < m {
		g.min.PushBack(data)
	} else {
		g.min.PushBack(m)
	}
}

func (g *GetMinStack) Pop() (int, bool) {
	e := g.data.Back()
	if e == nil {
		return 0, false
	}
	g.data.Remove(e)
	g.min.Remove(g.min.Back())
	return e.Value.(int), true
}

func (g *GetMinStack) GetMin() (int, bool) {
	e := g.min.Back()
	if e == nil {
		return 0, false
	}
	return e.Value.(int), true
}
