package winter06

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

// 排序链表
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}

	dummyNode := &ListNode{
		Val:  -1,
		Next: head,
	}

	for mergeUnit := 1; mergeUnit < length; mergeUnit <<= 1 {
		pre, cur := dummyNode, dummyNode.Next
		for cur != nil {
			var head1, head2 *ListNode

			head1 = cur
			for i := 1; i < mergeUnit && cur.Next != nil; i++ {
				cur = cur.Next
			}

			head2 = cur.Next
			cur.Next = nil
			cur = head2
			var next *ListNode
			for i := 1; i < mergeUnit && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			pre.Next = merge(head1, head2)

			for pre.Next != nil {
				pre = pre.Next
			}
			cur = next
		}
	}
	return dummyNode.Next
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	cur, p1, p2 := dummyNode, head1, head2
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

	if p1 == nil {
		cur.Next = p2
	}
	if p2 == nil {
		cur.Next = p1
	}
	return dummyNode.Next
}
