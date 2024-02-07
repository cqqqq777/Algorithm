package winter21

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack []*ListNode

func (s *Stack) Push(node *ListNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *ListNode {
	if len(*s) == 0 {
		return nil
	}
	node := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return node
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := Stack{}, Stack{}
	for cur := l1; cur != nil; cur = cur.Next {
		s1.Push(cur)
	}
	for cur := l2; cur != nil; cur = cur.Next {
		s2.Push(cur)
	}

	var carry int
	var p1, p2 *ListNode
	flag := len(s1) >= len(s2)
	for len(s1) > 0 || len(s2) > 0 || carry != 0 {
		p1, p2 = s1.Pop(), s2.Pop()
		if p1 != nil || p2 != nil || carry != 0 {
			cur := carry
			if p1 != nil {
				cur += p1.Val
			}
			if p2 != nil {
				cur += p2.Val
			}
			carry = cur / 10
			if p1 == nil && p2 == nil {
				if flag {
					p1 = &ListNode{Val: cur % 10, Next: l1}
					l1 = p1
				} else {
					p2 = &ListNode{Val: cur % 10, Next: l2}
					l2 = p2
				}
				break
			}
			if flag {
				p1.Val = cur % 10
			} else {
				p2.Val = cur % 10
			}
		}
	}
	if flag {
		return l1
	}
	return l2
}
