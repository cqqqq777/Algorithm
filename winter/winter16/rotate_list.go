package winter16

// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	// 找到倒数第 k 个节点，把后面的拼接到链表头部即可
	slow, fast := head, head
	for i := 0; i < k; i++ {
		// 有可能出现 k 大于链表长度，故出现 k 为 nil 的时候，需要回到链表头部
		// 先得出链表长度后对 k 取模也行
		fast = fast.Next
		if fast == nil {
			fast = head
			continue
		}
	}
	// k 等于链表长度
	if fast == head {
		return head
	}
	var pre, tail *ListNode
	for fast != nil {
		tail = fast
		fast = fast.Next
		pre = slow
		slow = slow.Next
	}

	pre.Next = nil
	tail.Next = head
	return slow
}
