package graph

import "testing"

func TestGenerateGraph(t *testing.T) {
	GenerateGraph([][]int{{1, 1, 2}, {1, 2, 3}, {1, 3, 2}})
}
