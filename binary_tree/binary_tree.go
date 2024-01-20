package binary_tree

type Node struct {
	Data        int
	Left, Right *Node
}

func NewBinaryTreeNode() *Node {
	return &Node{
		Left:  new(Node),
		Right: new(Node),
	}
}

func Find2ndNode(root *Node) *Node {
	if root == nil {
		return nil
	}

	cur, parent := root, NewBinaryTreeNode()
	for cur.Right != nil {
		parent = cur
		cur = cur.Right
	}

	// if rightest node's left != nil
	// the 2nd node is node's left tree's rightest node
	if cur.Left != nil {
		cur = cur.Left
		for cur.Right != nil {
			cur = cur.Right
		}
		return cur
	}

	return parent
}
