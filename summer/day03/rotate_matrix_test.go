package day03

import (
	"fmt"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	printMatrix(matrix)
	RotateMatrix(matrix)
	printMatrix(matrix)
}

func printMatrix(m [][]int) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			fmt.Printf("%d   ", m[i][j])
		}
		fmt.Println()
	}
}
