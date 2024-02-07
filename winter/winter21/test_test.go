package winter21

import (
	"testing"
)

func Test(t *testing.T) {
	l1, l2 := &ListNode{
		Val:  5,
		Next: nil,
	}, &ListNode{
		Val:  5,
		Next: nil,
	}
	addTwoNumbers(l1, l2)
}
