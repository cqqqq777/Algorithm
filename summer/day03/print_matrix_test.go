package day03

import "testing"

func TestSpiralPrint(t *testing.T) {
	SpiralPrint([][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	})
}

func TestZigZagPrint(t *testing.T) {
	ZigZagPrint([][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	})
}
