package graph

import "fmt"

// BFS 宽度优先遍历
func BFS(head *Node) {
	if head == nil {
		return
	}
	queue := new(Queue)
	// map used to store nodes which have been queued
	nodes := make(map[*Node]struct{})
	queue.Add(head)
	nodes[head] = struct{}{}
	for !queue.IsEmpty() {
		cur := queue.Poll()
		fmt.Println(cur.Value)
		for _, v := range cur.Nexts {
			if _, ok := nodes[v]; !ok {
				queue.Add(v)
				nodes[v] = struct{}{}
			}
		}
	}
}

// DFS 深度优先遍历
func DFS(head *Node) {
	if head == nil {
		return
	}
	stack := new(Stack)
	nodes := make(map[*Node]struct{})
	stack.Push(head)
	nodes[head] = struct{}{}
	fmt.Println(head.Value)
	for !stack.IsEmpty() {
		cur := stack.Pop()
		for _, v := range cur.Nexts {
			if _, ok := nodes[v]; !ok {
				stack.Push(cur)
				stack.Push(v)
				nodes[v] = struct{}{}
				fmt.Println(v.Value)
				break
			}
		}
	}
}
