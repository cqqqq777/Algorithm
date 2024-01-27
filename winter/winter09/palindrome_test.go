package winter09

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	fmt.Println(isPalindrome(head))
}
