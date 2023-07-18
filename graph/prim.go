package graph

import (
	"container/heap"
	"fmt"
)

type PriorityQueue []*Edge

func (p PriorityQueue) Len() int { return len(p) }

func (p PriorityQueue) Less(i, j int) bool { return p[i].Weight < p[j].Weight }

func (p PriorityQueue) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p *PriorityQueue) Push(x interface{}) {
	edge := x.(*Edge)
	*p = append(*p, edge)
}

func (p *PriorityQueue) Pop() interface{} {
	old := *p
	*p = old[0 : len(old)-1]
	return old[len(old)-1]
}

// Prim 生成最小生成树
func Prim(graph *Graph) []*Edge {
	// priorityQueue used to store unlock edge
	priorityQueue := make(PriorityQueue, 0)
	heap.Init(&priorityQueue)
	sets := make(map[*Node]struct{})
	result := make([]*Edge, 0)
	for _, v := range graph.Nodes {
		if _, ok := sets[v]; !ok {
			sets[v] = struct{}{}
			for _, edge := range v.Edges {
				heap.Push(&priorityQueue, edge)
			}
			for len(priorityQueue) != 0 {
				edge := heap.Pop(&priorityQueue).(*Edge)
				node := edge.To
				if _, exist := sets[node]; !exist {
					sets[node] = struct{}{}
					result = append(result, edge)
					for _, nextEdge := range node.Edges {
						heap.Push(&priorityQueue, nextEdge)
					}
				}
			}
		}
	}
	return result
}

func Priority() {
	priorityQueue := make(PriorityQueue, 0)
	//for i := 0; i < 5; i++ {
	//	edge := &Edge{Weight: i}
	//	priorityQueue = append(priorityQueue, edge)
	//}
	heap.Init(&priorityQueue)
	//fmt.Println(heap.Pop(&priorityQueue).(*Edge).Weight)
	heap.Push(&priorityQueue, &Edge{Weight: 9})
	heap.Push(&priorityQueue, &Edge{Weight: 8})
	heap.Push(&priorityQueue, &Edge{Weight: 3})
	fmt.Println(heap.Pop(&priorityQueue).(*Edge).Weight)
	fmt.Println(len(priorityQueue))
}
