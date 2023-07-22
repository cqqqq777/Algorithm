package dequeue

import (
	"fmt"
	"testing"
)

func TestDequeue(t *testing.T) {
	dq := NewDequeue()
	dq.AddAtTail(1)
	dq.AddAtHead(2)
	dq.AddAtTail(3)
	dq.AddAtHead(4)
	fmt.Println(dq.Size())
	fmt.Println(dq.IsEmpty())
	fmt.Println(dq.PoolAtHead())
	fmt.Println(dq.PoolAtHead())
	fmt.Println(dq.Size())
	fmt.Println(dq.IsEmpty())
}
