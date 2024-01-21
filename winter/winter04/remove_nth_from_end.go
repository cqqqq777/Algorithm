package winter04

// 删除倒数第 n 个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}

	m, i := make(map[int]*ListNode), 1
	for cur := head; cur != nil; cur, i = cur.Next, i+1 {
		m[i] = cur
	}

	node, ok := m[i-n-1]
	if !ok {
		return head
	}
	node.Next = node.Next.Next
	return head
}

// 双指针
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}

	dummyNode := &ListNode{
		Val:  -1,
		Next: head,
	}
	first, second := dummyNode, head
	// second 先走 n 步，那还剩 k - n 步
	for i := 0; i < n; i++ {
		second = second.Next
	}
	// first 再跟着 second 走 k - n 步，first 就来到倒数第 n-1 个节点
	for second != nil {
		first, second = first.Next, second.Next
	}
	first.Next = first.Next.Next
	return dummyNode.Next
}
