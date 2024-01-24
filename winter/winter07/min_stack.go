package winter07

type MinStack struct {
	s1, s2 []int       // 一个正常栈一个单调栈
	m      map[int]int // 记录栈中元素的频次
}

func Constructor() MinStack {
	return MinStack{
		s1: make([]int, 0),
		s2: make([]int, 0),
		m:  make(map[int]int),
	}
}

func (m *MinStack) Push(val int) {
	m.m[val]++
	m.s1 = append(m.s1, val)
	count := 0
	for len(m.s2) > 0 && m.s2[len(m.s2)-1] < val {
		m.s1 = append(m.s1, m.s2[len(m.s2)-1])
		m.s2 = m.s2[:len(m.s2)-1]
		count++
	}
	m.s2 = append(m.s2, val)
	for i := 0; i < count; i++ {
		m.s2 = append(m.s2, m.s1[len(m.s1)-1])
		m.s1 = m.s1[:len(m.s1)-1]
	}
}

func (m *MinStack) Pop() {
	val := m.s1[len(m.s1)-1]
	m.m[val]--
	m.s1 = m.s1[:len(m.s1)-1]
}

func (m *MinStack) Top() int {
	return m.s1[len(m.s1)-1]
}

func (m *MinStack) GetMin() int {
	for {
		val := m.s2[len(m.s2)-1]
		if m.m[val] > 0 {
			return val
		} else {
			m.s2 = m.s2[:len(m.s2)-1]
		}
	}
}
