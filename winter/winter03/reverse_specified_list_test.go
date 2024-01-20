package winter03

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	//head := &ListNode{}
	//cur := head
	//for i := 1; i <= 4; i++ {
	//	cur.Val = i
	//	cur.Next = &ListNode{}
	//	cur = cur.Next
	//}
	//cur.Val = 5
	//head = reverseBetween(head, 2, 4)
	//for cur := head; cur != nil; cur = cur.Next {
	//	fmt.Println(cur.Val)
	//}
	var Nil interface{}
	_, ok := Nil.(*ListNode)
	fmt.Println(ok)
}
