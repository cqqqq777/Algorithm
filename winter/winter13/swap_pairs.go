package winter13

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{
		Next: head,
	}
	pre, cur, next := dummy, head, head.Next
	for cur != nil && next != nil {
		cur.Next = next.Next
		next.Next = cur
		pre.Next = cur
		pre = cur
		if cur.Next == nil || cur.Next.Next == nil {
			break
		}
		cur = cur.Next
		next = cur.Next
	}
	return dummy.Next
}
