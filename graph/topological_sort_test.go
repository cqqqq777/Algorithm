package graph

import (
	"fmt"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	graph := GenerateGraph([][]int{{1, 1, 2}, {1, 1, 3}, {1, 2, 4}, {1, 4, 3}, {1, 1, 4}})
	result := TopologicalSort(graph)
	for _, v := range result {
		fmt.Println(v.Value)
	}
}
