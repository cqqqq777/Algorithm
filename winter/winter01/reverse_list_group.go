package winter01

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	return process(head, k)
}

func process(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	// 判断剩余的是否够 k 个一组进行翻转
	end := head
	for i := 0; i < k-1; i++ {
		end = end.Next
		if end == nil {
			return head
		}
	}

	// 翻转 k 次
	cur := head
	var pre, next *ListNode
	for i := 0; i < k; i++ {
		next = cur.Next
		if next == nil {
			return head
		}
		cur.Next = pre
		pre, cur = cur, next
	}
	// 此时 pre 为翻转后的头节点，head 为翻转后的尾节点
	nextHead := process(next, k)
	// 尾节点与下一组的头节点相连
	head.Next = nextHead
	// 返回头节点
	return pre
}
