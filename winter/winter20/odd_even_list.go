package winter20

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。
// 第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。
// 请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。
// 你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	head1 := head.Next
	odd, even := head, head1
	for {
		if odd.Next == nil {
			break
		}
		oddNext := odd.Next.Next
		if odd.Next.Next != nil {
			odd.Next = oddNext
			odd = oddNext
		}

		if even.Next == nil {
			break
		}
		evenNext := even.Next.Next
		even.Next = evenNext
		even = evenNext
	}
	odd.Next = head1
	return head
}
