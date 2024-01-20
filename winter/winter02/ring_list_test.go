package winter02

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	fmt.Println(hasCycle(h))
}
