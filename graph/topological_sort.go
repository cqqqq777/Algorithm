package graph

// TopologicalSort 拓扑排序
func TopologicalSort(graph *Graph) []*Node {
	if graph == nil {
		return []*Node{}
	}

	result := make([]*Node, len(graph.Nodes))
	inMap := make(map[*Node]int)

	i := 0
	for _, v := range graph.Nodes {
		inMap[v] = v.In
		if v.In == 0 {
			result[i] = v
			i++
		}
	}

	for cur := 0; cur != len(graph.Nodes); cur++ {
		curNode := result[cur]
		for _, v := range curNode.Nexts {
			inMap[v] = inMap[v] - 1
			if inMap[v] == 0 {
				result[i] = v
				i++
			}
		}
	}

	return result
}
