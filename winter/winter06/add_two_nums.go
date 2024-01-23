package winter06

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}

	p1, p2 := l1, l2
	// 记录进位信息
	carry, cur := 0, dummyNode
	for ; p1 != nil && p2 != nil; p1, p2 = p1.Next, p2.Next {
		res := p1.Val + p2.Val + carry
		cur.Next = &ListNode{}
		cur.Next.Val = res % 10
		cur = cur.Next
		carry = res / 10
	}
	for ; p1 != nil; p1 = p1.Next {
		res := p1.Val + carry
		carry = res / 10
		cur.Next = &ListNode{}
		cur.Next.Val = res % 10
		cur = cur.Next
	}
	for ; p2 != nil; p2 = p2.Next {
		res := p2.Val + carry
		carry = res / 10
		cur.Next = &ListNode{}
		cur.Next.Val = res % 10
		cur = cur.Next
	}
	if carry != 0 {
		cur.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return dummyNode.Next
}
