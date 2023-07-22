package list

// ReverseList 反转链表
func ReverseList(head *Node) *Node {
	var pre *Node
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// Reverse 递归方式反转一个链表
func Reverse(pre, cur *Node) *Node {
	if cur == nil {
		return pre
	}
	next := cur.Next
	cur.Next = pre
	pre = cur
	return Reverse(pre, next)
}

type Stack struct{}

func (s *Stack) Add(_ *Node) {}

func (s *Stack) Pop() *Node {
	return nil
}

func (s *Stack) IsEmpty() bool {
	return true
}

// ReverseL 栈（以下虚拟出一个栈，并没有实现栈）
func ReverseL(head *Node) *Node {
	stack := new(Stack)

	cur := head
	for cur != nil {
		stack.Add(cur)
		cur = cur.Next
	}

	newHead := stack.Pop()
	cur = newHead

	for !stack.IsEmpty() {
		cur.Next = stack.Pop()
		cur = cur.Next
	}

	cur.Next = nil

	return newHead
}

// ReverseDoubleList 反转双向链表
func ReverseDoubleList(head *DoubleListNode) *DoubleListNode {
	cur := head

	for cur != nil {
		cur.Next, cur.Last = cur.Last, cur.Next
		head = cur
		cur = cur.Last
	}

	return head
}

// FindEqualNode 找出两个有序链表相等的部分
func FindEqualNode(head1, head2 *Node) []int {
	result := make([]int, 0)

	p1, p2 := head1, head2

	for p1 != nil && p2 != nil {
		if p1.Data > p2.Data {
			p2 = p2.Next
		} else if p1.Data < p2.Data {
			p1 = p1.Next
		} else {
			result = append(result, p1.Data)
			p1 = p1.Next
			p2 = p2.Next
		}
	}

	return result
}

// PalindromeList1 判断一个链表是否为回文链表(使用栈)
func PalindromeList1(head *Node) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow = slow.Next

	stack := new(Stack)
	for slow != nil {
		stack.Add(slow)
		slow = slow.Next
	}

	slow = head
	for !stack.IsEmpty() {
		if slow.Data != stack.Pop().Data {
			return false
		}
	}

	return true
}

// PalindromeList2 判断一个链表是否为回文链表(空间复杂度为O(1))
func PalindromeList2(head *Node) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var pre *Node
	cur := slow
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	for head != nil {
		if head.Data != pre.Data {
			return false
		}
		head = head.Next
		pre = pre.Next
	}

	return true
}

// CopyRandList1 复制一个含有随机指针的链表(使用 map，额外空间复杂度为O(N))
func CopyRandList1(head *RandNode) *RandNode {
	if head == nil {
		return nil
	}

	kv := make(map[*RandNode]*RandNode)

	for cur := head; cur != nil; cur = cur.Next {
		newNode := new(RandNode)
		newNode.Data = cur.Data
		kv[cur] = newNode
	}

	for cur := head; cur != nil; cur = cur.Next {
		kv[cur].Next = kv[cur.Next]
		kv[cur].Rand = kv[cur.Rand]
	}

	return kv[head]
}

// CopyRandList2 复制一个含有随机指针的链表(不适用额外数据结构，额外空间复杂度O(1))
func CopyRandList2(head *RandNode) *RandNode {
	if head == nil {
		return nil
	}

	// copy old nodes and put them in the back of old nodes
	for cur := head; cur != nil; cur = cur.Next {
		newNode := new(RandNode)
		newNode.Data = cur.Data
		newNode.Next = cur.Next
		cur.Next = newNode
	}

	// set the rand
	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next.Rand = cur.Rand.Next
	}

	// district old nodes and new nodes
	for cur := head; cur != nil; {
		next := cur.Next.Next
		cur.Next.Next = cur.Next.Next.Next
		cur = next
	}

	return head.Next
}

// CheckLoop1 使用 map 法，时间复杂度和额外空间复杂度都是 O(N)
func CheckLoop1(head *Node) (*Node, bool) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil, false
	}

	kv := make(map[*Node]struct{})
	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := kv[cur]; ok {
			return cur, true
		}
	}
	return nil, false
}

// CheckLoop2 快慢指针法，额外空间复杂度为O(1)
func CheckLoop2(head *Node) (*Node, bool) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil, false
	}

	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast, true
		}
	}

	return nil, false
}

// FindFirstIntersectNode 判断两个链表是否有相交部分并返回相交点
func FindFirstIntersectNode(head1, head2 *Node) *Node {
	loop1, ok1 := CheckLoop2(head1)
	_, ok2 := CheckLoop2(head2)

	if !ok1 && !ok2 {
		// get their length
		len1, len2 := 0, 0
		for cur := head1; cur != nil; cur = cur.Next {
			len1++
		}
		for cur := head2; cur != nil; cur = cur.Next {
			len2++
		}

		// the longer list move |n| times first
		n := len1 - len2
		p1, p2 := head1, head2
		if n > 0 {
			for i := 0; i < n; i++ {
				p1 = p1.Next
			}
		} else if n < 0 {
			for i := 0; i < -n; i++ {
				p2 = p2.Next
			}
		}

		// p1,p2 move together
		for p1 != p2 {
			p1 = p1.Next
			p2 = p2.Next
		}

		return p1
	} else if ok1 && ok2 {
		return loop1
	}

	return nil
}
