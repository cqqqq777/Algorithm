package winter01

import "testing"

func TestReverseListGroup(t *testing.T) {
	head := &ListNode{}
	cur := head
	for i := 1; i <= 4; i++ {
		cur.Val = i
		cur.Next = &ListNode{}
		cur = cur.Next
	}
	cur.Val = 5
	reverseKGroup(head, 2)
}
