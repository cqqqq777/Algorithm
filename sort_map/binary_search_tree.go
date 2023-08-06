package sort_map

type Key int64

type Value int64

type Node struct {
	Key                 Key
	Value               Value
	Left, Right, Parent *Node
}

type BinarySearchTree struct {
	Root *Node
	Size int
}

// NewBinarySearchTree returns a new BinarySearchTree.
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (bst *BinarySearchTree) Insert(key Key, value Value) {
	if bst.Root == nil {
		bst.Root = &Node{Key: key, Value: value}
		bst.Size++
		return
	}
	bst.insert(bst.Root, key, value)
}

func (bst *BinarySearchTree) insert(node *Node, key Key, value Value) {
	cur := node
	for {
		if cur.Key > key {
			if cur.Left == nil {
				cur.Left = &Node{Key: key, Value: value, Parent: cur}
				bst.Size++
				return
			}
			cur = cur.Left
		} else if cur.Key < key {
			if cur.Right == nil {
				cur.Right = &Node{Key: key, Value: value, Parent: cur}
				bst.Size++
				return
			}
			cur = cur.Right
		} else {
			cur.Value = value
			return
		}
	}
}

// Search returns the value of the key.
func (bst *BinarySearchTree) Search(key Key) (Value, bool) {
	return bst.search(bst.Root, key)
}

func (bst *BinarySearchTree) search(node *Node, key Key) (Value, bool) {
	cur := node
	for cur.Key != key && cur != nil {
		if cur.Key > key {
			cur = cur.Left
		}
		if cur.Key < key {
			cur = cur.Right
		}
	}
	if cur == nil {
		return 0, false
	}
	return cur.Value, true
}

// Delete deletes the key.
func (bst *BinarySearchTree) Delete(key Key) {
	bst.delete(bst.Root, key)
}

func (bst *BinarySearchTree) delete(node *Node, key Key) {
	// search key first
	cur := node
	for cur != nil && cur.Key != key {
		if cur.Key > key {
			cur = cur.Left
		}
		if cur.Key < key {
			cur = cur.Right
		}
	}
	if cur == nil {
		return
	}
	// delete node
	if cur.Left == nil && cur.Right == nil {
		if cur.Parent == nil {
			bst.Root = nil
			bst.Size--
			return
		}
		if cur.Parent.Left == cur {
			cur.Parent.Left = nil
			bst.Size--
			return
		} else {
			cur.Parent.Right = nil
			bst.Size--
			return
		}
	}
	if cur.Left == nil {
		if cur.Parent == nil {
			bst.Root = nil
			bst.Size--
			return
		}
		if cur.Parent.Left == cur {
			cur.Parent.Left = cur.Right
			cur.Right.Parent = cur.Parent
			bst.Size--
			return
		} else {
			cur.Parent.Right = cur.Right
			cur.Right.Parent = cur.Parent
			bst.Size--
			return
		}
	}
	if cur.Right == nil {
		if cur.Parent == nil {
			bst.Root = nil
			bst.Size--
			return
		}
		if cur.Parent.Left == cur {
			cur.Parent.Left = cur.Left
			cur.Left.Parent = cur.Parent
			bst.Size--
			return
		} else {
			cur.Parent.Right = cur.Left
			cur.Left.Parent = cur.Parent
			bst.Size--
			return
		}
	}
	// find successor
	successor := cur.Right
	for successor.Left != nil {
		successor = successor.Left
	}
	cur.Key = successor.Key
	cur.Value = successor.Value
	// delete successor
	if successor.Parent.Left == successor {
		if successor.Right != nil {
			successor.Right.Parent = successor.Parent
		}
		successor.Parent.Left = successor.Right
	} else {
		if successor.Right != nil {
			successor.Right.Parent = successor.Parent
		}
		successor.Parent.Right = successor.Right
	}
	bst.Size--
}
