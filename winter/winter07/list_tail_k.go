package winter07

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给定一个头节点为 head 的链表用于记录一系列核心肌群训练项目编号，请查找并返回倒数第 cnt 个训练项目编号。
func trainingPlan(head *ListNode, cnt int) *ListNode {
	first, second := head, head
	for i := 0; i < cnt; i++ {
		second = second.Next
	}
	for second != nil {
		second = second.Next
		first = first.Next
	}
	return first
}
