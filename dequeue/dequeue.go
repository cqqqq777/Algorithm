package dequeue

type Dequeue struct {
	head, tail *Node
}

type Node struct {
	value          int
	next, previous *Node
}

func NewDequeue() *Dequeue {
	return &Dequeue{}
}

func (d *Dequeue) AddAtHead(value int) {
	node := &Node{value: value}
	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		node.next = d.head
		d.head.previous = node
		d.head = node
	}
}

func (d *Dequeue) AddAtTail(value int) {
	node := &Node{value: value}
	if d.tail == nil {
		d.head = node
		d.tail = node
	} else {
		d.tail.next = node
		node.previous = d.tail
		d.tail = node
	}
}

func (d *Dequeue) PoolAtHead() int {
	value := d.head.value
	if d.head.next == nil {
		d.tail = nil
		d.head = d.head.next
	} else {
		d.head = d.head.next
		d.head.previous = nil
	}
	return value
}

func (d *Dequeue) PoolAtTail() int {
	value := d.tail.value
	if d.tail.previous == nil {
		d.head = nil
		d.tail = d.tail.previous
	} else {
		d.tail = d.tail.previous
		d.tail.next = nil
	}
	return value
}

func (d *Dequeue) PeekAtHead() int {
	return d.head.value
}

func (d *Dequeue) PeekAtTail() int {
	return d.tail.value
}

func (d *Dequeue) Size() int {
	if d.head == nil {
		return 0
	}
	size := 1
	head, tail := d.head, d.tail
	for ; head != tail; head = head.next {
		size++
	}
	return size
}

func (d *Dequeue) IsEmpty() bool {
	return d.head == nil
}
