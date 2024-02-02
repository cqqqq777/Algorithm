package winter16

// 给定一个奇数位升序，偶数位降序的链表，将其重新排序。
// 输入: 1->8->3->6->5->4->7->2->NULL
// 输出: 1->2->3->4->5->6->7->8->NULL

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head1, head2 := head, head.Next
	p1, p2 := head1, head2
	// 拆成两个链表
	for {
		next1 := p1.Next.Next
		if next1 == nil {
			p1.Next = nil
			break
		}
		p1.Next = next1

		next2 := p2.Next.Next
		if next2 == nil {
			p2.Next = nil
			break
		}
		p2.Next = next2

		p1, p2 = next1, next2
	}

	// 逆序降序链表
	p2 = reverseList(head2)
	p1 = head1

	// 合并两个有序链表
	dummy := &ListNode{
		Next: head,
	}
	cur := dummy
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			cur.Next = p2
			p2 = p2.Next
		} else {
			cur.Next = p1
			p1 = p1.Next
		}
		cur = cur.Next
	}
	if p1 != nil {
		cur.Next = p1
	} else {
		cur.Next = p2
	}
	return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
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
