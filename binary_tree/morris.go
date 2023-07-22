package binary_tree

import "fmt"

// Morris 遍历二叉树
func Morris(head *Node) {
	cur := head
	for cur != nil {
		if cur.Left == nil {
			cur = cur.Right // 如果左节点为空就往右走
		} else {
			mostRight := cur.Left // 找到左子树上的最右节点
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur // 用最右节点的右指针标记 cur 用于下次跳回 cur
				cur = cur.Left        // cur 往左走
			} else {
				mostRight.Right = nil // 此时 cur 已经被标记过了，将最右节点的右指针还原
				cur = cur.Right       // cur 往右走
			}
		}
	}
}

// MorrisPre 先序
func MorrisPre(head *Node) {
	cur := head
	for cur != nil {
		if cur.Left == nil {
			fmt.Println(cur.Data) // 没有左子树的节点只会到达一次
			cur = cur.Right       // 如果左节点为空就往右走
		} else {
			mostRight := cur.Left // 找到左子树上的最右节点
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur // 用最右节点的右指针标记 cur 用于下次跳回 cur
				fmt.Println(cur.Data) // 此时是第一次来到 cur，处理
				cur = cur.Left        // cur 往左走
			} else {
				mostRight.Right = nil // 此时 cur 已经被标记过了，将最右节点的右指针还原
				cur = cur.Right       // cur 往右走
			}
		}
	}
}

// MorrisIn 中序
func MorrisIn(head *Node) {
	cur := head
	for cur != nil {
		if cur.Left == nil {
			fmt.Println(cur.Data)
			cur = cur.Right // 如果左节点为空就往右走
		} else {
			mostRight := cur.Left // 找到左子树上的最右节点
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur // 用最右节点的右指针标记 cur 用于下次跳回 cur
				cur = cur.Left        // cur 往左走
			} else {
				mostRight.Right = nil // 此时 cur 已经被标记过了，将最右节点的右指针还原
				fmt.Println(cur.Data) // 这是第二次来到 cur，处理 cur
				cur = cur.Right       // cur 往右走
			}
		}
	}
}

// MorrisPos 后序
func MorrisPos(head *Node) {
	cur := head
	for cur != nil {
		if cur.Left == nil {
			cur = cur.Right // 如果左节点为空就往右走
		} else {
			mostRight := cur.Left // 找到左子树上的最右节点
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur // 用最右节点的右指针标记 cur 用于下次跳回 cur
				cur = cur.Left        // cur 往左走
			} else {
				mostRight.Right = nil // 此时 cur 已经被标记过了，将最右节点的右指针还原
				reverse(cur.Left)
				for node := mostRight; node != nil; node = node.Right {
					fmt.Println(node.Data)
				}
				reverse(mostRight)
				cur = cur.Right // cur 往右走
			}
		}
	}
	h := reverse(head)
	for node := h; node != nil; node = node.Right {
		fmt.Println(node.Data)
	}
	reverse(h)
}

func reverse(node *Node) *Node {
	var pre *Node
	cur := node
	for cur != nil {
		next := cur.Right
		cur.Right = pre
		pre = cur
		cur = next
	}
	return pre
}
