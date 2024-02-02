package winter16

import (
	"fmt"
	"testing"
)

// 输入: 1->8->3->6->5->4->7->2->NULL
// 输出: 1->2->3->4->5->6->7->8->NULL
func Test(t *testing.T) {
	//head := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 8,
	//		Next: &ListNode{
	//			Val: 3,
	//			Next: &ListNode{
	//				Val: 6,
	//				Next: &ListNode{
	//					Val: 5,
	//					Next: &ListNode{
	//						Val: 4,
	//						Next: &ListNode{
	//							Val: 7,
	//							Next: &ListNode{
	//								Val:  2,
	//								Next: nil,
	//							},
	//						},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	//head = sortList(head)
	//for cur := head; cur != nil; cur = cur.Next {
	//	fmt.Println(cur.Val)
	//}
	fmt.Println(reversePairs([]int{}))
}
