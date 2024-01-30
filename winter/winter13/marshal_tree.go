package winter13

import (
	"fmt"
	"strconv"
	"strings"
)

// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "*!"
	}
	left, right := c.serialize(root.Left), c.serialize(root.Right)
	return fmt.Sprintf("%d!%s%s", root.Val, left, right)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	values := strings.Split(data, "!")
	values = values[:len(values)-1]
	queue := make([]string, 0)
	for _, v := range values {
		queue = append(queue, v)
	}
	return buildByPreOrder(&queue)
}

func buildByPreOrder(queue *[]string) *TreeNode {
	if len(*queue) == 0 {
		return nil
	}
	val := (*queue)[0]
	*queue = (*queue)[1:]
	if val == "*" {
		return nil
	}
	num, _ := strconv.Atoi(val)
	head := &TreeNode{
		Val: num,
	}
	head.Left = buildByPreOrder(queue)
	head.Right = buildByPreOrder(queue)
	return head
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
