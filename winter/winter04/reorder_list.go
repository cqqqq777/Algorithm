package winter04

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

// 将链表重新排列成 L0 -> Ln -> L1 -> Ln-1 -> L2 -> Ln-2 ...
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	// find middle node
	p1, p2 := head, head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	// reverse post
	p2 = reverseList(p1.Next)
	p1.Next = nil
	p1 = head

	// insert node
	for p1 != nil && p2 != nil {
		next1, next2 := p1.Next, p2.Next
		p1.Next = p2
		p2.Next = next1
		p1, p2 = next1, next2
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
