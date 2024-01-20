package graph

import (
	"fmt"
	"testing"
)

func TestPriority(t *testing.T) {
	Priority()
}

func TestPrim(t *testing.T) {
	graph := GenerateGraph([][]int{{2, 1, 3}, {7, 1, 2}, {100, 1, 4}, {4, 3, 4}, {1000, 2, 4}, {10000, 2, 5}})
	result := Prim(graph)
	for _, v := range result {
		fmt.Printf("%d ", v.Weight) // expected output: 2 4 7 10000
	}
}
