package winter03

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	n1, n2 := &ListNode{
		Val:  1,
		Next: nil,
	}, &ListNode{
		Val:  2,
		Next: nil,
	}
	pq := make(PriorityQueue, 2)
	pq[0] = n2
	pq[1] = n1
	heap.Init(&pq)
	fmt.Println(pq.Pop().(*ListNode).Val)
	fmt.Println(pq.Pop().(*ListNode).Val)

}
