package winter09

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast.Next != nil {
		fast = fast.Next
		next := slow.Next
		slow.Next = nil
		reverseList(next, nil)
	} else {
		next := slow.Next
		slow.Next = nil
		reverseList(next, slow)
	}

	slow = head
	for ; slow != nil && fast != nil; slow, fast = slow.Next, fast.Next {
		if slow.Val != fast.Val {
			return false
		}
	}
	return true
}

func reverseList(head, pre *ListNode) *ListNode {
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
