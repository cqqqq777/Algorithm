package list

type Node struct {
	Data int
	Next *Node
}

type DoubleListNode struct {
	Data       int
	Next, Last *DoubleListNode
}

type RandNode struct {
	Data       int
	Next, Rand *RandNode
}

func AddNode(header **Node, data, position int) {
	cur := *header

	newNode := new(Node)
	newNode.Data = data

	if cur == nil {
		cur = newNode
		return
	}

	for i := 0; i < position; i++ {
		if cur.Next == nil {
			cur.Next = newNode
			return
		}
		cur = cur.Next
	}

	newNode.Next = cur.Next
	cur.Next = newNode
}

func DeleteNode(header **Node, position int) {
	cur := *header
	for i := 0; i < position-1; i++ {
		if cur.Next == nil {
			return
		}
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
}
