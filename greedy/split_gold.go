package greedy

import (
	"sort"
)

// SplitGold 要求用最少的钱达到给定的结果
func SplitGold(golds []int) int {
	sort.Ints(golds)
	sum := 0
	for len(golds) != 1 {
		cur := golds[0] + golds[1]
		golds[1] = cur
		golds = golds[1:]
		for i := 1; i < len(golds)-1; i++ {
			if golds[i-1] > golds[i] {
				golds[i-1], golds[i] = golds[i], golds[i-1]
			}
		}
		sum += cur
	}
	return sum
}
