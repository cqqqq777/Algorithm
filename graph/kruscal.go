package graph

import "sort"

// Kruscal 生成最小生成树
func Kruscal(graph *Graph) []*Edge {
	edges := make([]*Edge, len(graph.Edges))
	i := 0
	for k, _ := range graph.Edges {
		edges[i] = k
		i++
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})
	// map used to simulate sets
	// key is node, value is the number of set
	sets := make(map[*Node]int, 0)
	i = 1
	for _, v := range graph.Nodes {
		sets[v] = i
		i++
	}
	result := make([]*Edge, 0)
	for _, v := range edges {
		if sets[v.From] != sets[v.To] {
			// merge two sets
			// set the 'To' node's set to 'From' node's set
			sets[v.To] = sets[v.From]
			result = append(result, v)
		}
	}
	return result
}
