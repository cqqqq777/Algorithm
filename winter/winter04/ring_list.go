package winter04

// 找入环节点
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	if fast.Next == nil || fast.Next.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}
