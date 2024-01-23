package winter05

type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除有序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	dummyNode := &ListNode{
		Val:  -1,
		Next: head,
	}

	cur := dummyNode
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummyNode.Next
}
