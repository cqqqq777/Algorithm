package dequeue

// Stack 基于双端队列实现的栈
type Stack struct {
	d *Dequeue
}

func NewStack() *Stack {
	return &Stack{d: NewDequeue()}
}

func (s *Stack) Push(value int) {
	s.d.AddAtTail(value)
}
func (s *Stack) Pop() int {
	return s.d.PoolAtTail()
}

func (s *Stack) IsEmpty() bool {
	return s.d.IsEmpty()
}

func (s *Stack) Peek() int {
	return s.d.PeekAtTail()
}

// GetNearLess 获取数组每个元素左右两边距离最近且小于它的两个元素（给定数组无重复元素）
func GetNearLess(arr []int) [][]int {
	stack := NewStack()
	res := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		for !stack.IsEmpty() && arr[stack.Peek()] > arr[i] {
			cur := stack.Pop()
			if stack.IsEmpty() {
				res[cur] = []int{-1, i}
				break
			}
			res[cur] = []int{stack.Peek(), i}
		}
		stack.Push(i)
	}
	for !stack.IsEmpty() {
		cur := stack.Pop()
		if stack.IsEmpty() {
			res[cur] = []int{-1, -1}
			break
		}
		res[cur] = []int{stack.Peek(), -1}
	}
	return res
}

// GetMaxA 返回给定数组子数组的最大指标
func GetMaxA(arr []int) int {
	max := 0
	nearLess := GetNearLess(arr)
	for i := 0; i < len(arr); i++ {
		sum := 0
		left := nearLess[i][0] + 1
		right := nearLess[i][1]
		if right == -1 {
			right = len(arr)
		}
		for j := left; j < right; j++ {
			sum += arr[j]
		}
		if max < sum*arr[i] {
			max = sum * arr[i]
		}
	}
	return max
}
