package day03

import "fmt"

func ZigZagPrint(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	row1, col1, row2, col2, endRow, endCol := 0, 0, 0, 0, len(matrix), len(matrix[0])
	b := false
	for col1 < endCol {
		zigzagPrint(matrix, row1, col1, row2, col2, b)
		if row1 == endRow-1 {
			col1++
		} else {
			row1++
		}
		if col2 == endCol-1 {
			row2++
		} else {
			col2++
		}
		b = !b
	}
}

func zigzagPrint(m [][]int, row1, col1, row2, col2 int, b bool) {
	if b {
		for row2 <= row1 {
			fmt.Printf("%d ", m[row2][col2])
			col2--
			row2++
		}
	} else {
		for row1 >= row2 {
			fmt.Printf("%d ", m[row1][col1])
			col1++
			row1--
		}
	}
}

// SpiralPrint 螺旋打印矩阵
func SpiralPrint(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	r1, c1, r2, c2 := 0, 0, len(matrix)-1, len(matrix[0])-1
	for r1 <= r2 && c1 <= c2 {
		spiralPrint(matrix, r1, r2, c1, c2)
		r1++
		c1++
		r2--
		c2--
	}
}

func spiralPrint(m [][]int, row1, row2, col1, col2 int) {
	if row1 == row2 {
		for i := col1; i <= col2; i++ {
			fmt.Printf("%d ", m[row1][i])
		}
	} else if col1 == col2 {
		for i := row1; i <= row2; i++ {
			fmt.Printf("%d ", m[i][col1])
		}
	} else {
		for i := col1; i < col2; i++ {
			fmt.Printf("%d ", m[row1][i])
		}
		for i := row1; i < row2; i++ {
			fmt.Printf("%d ", m[i][col2])
		}
		for i := col2; i > col1; i-- {
			fmt.Printf("%d ", m[row2][i])
		}
		for i := row2; i > row1; i-- {
			fmt.Printf("%d ", m[i][col1])
		}
	}
}
