package winter03

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
// 记录 left 前一个节点 pre，每次插一个节点到 pre 和 left 之间
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	// 虚拟节点
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	// pre 永远为 left 的前一个节点
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := 0; i < right-left; i++ {
		// 每次都将节点插到 pre 的下一个
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return dummyNode.Next
}
