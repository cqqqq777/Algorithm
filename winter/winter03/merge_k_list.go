package winter03

import (
	"container/heap"
)

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Val < pq[j].Val }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// 合并 k 个有序数组
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	pq := make(PriorityQueue, 0, len(lists))
	for _, v := range lists {
		if v != nil {
			pq = append(pq, v)
		}
	}
	heap.Init(&pq)

	dummyNode := &ListNode{Val: -1}
	cur := dummyNode
	for len(pq) != 0 {
		minNode := heap.Pop(&pq).(*ListNode)
		cur.Next = minNode
		cur = minNode
		if minNode.Next != nil {
			heap.Push(&pq, minNode.Next)
		}
	}

	return dummyNode.Next
}

// 递归，每次合并两个链表
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	l := len(lists) / 2
	pre := mergeKLists(lists[:l])
	post := mergeKLists(lists[l:])
	return mergeList(pre, post)
}

func mergeList(l1, l2 *ListNode) *ListNode {
	var dummyHead = ListNode{
		Val:  0,
		Next: nil,
	}
	cur := &dummyHead
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				cur.Next = l1
				l1 = l1.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next
		} else if l1 != nil {
			cur.Next = l1
			break
		} else if l2 != nil {
			cur.Next = l2
			break
		}
	}
	return dummyHead.Next
}
