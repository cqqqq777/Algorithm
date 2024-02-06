package winter20

import (
	"math"
	"sort"
)

// 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
// 返回这三个数的和。
// 假定每组输入只存在恰好一个解。
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	res := math.MaxInt
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		p1, p2 := i+1, len(nums)-1
		for p1 < p2 {
			cur := nums[i] + nums[p1] + nums[p2]
			if cur == target {
				return target
			}
			res = choose(res, cur, target)

			if cur > target {
				p := p2 - 1
				for p1 < p && nums[p] == nums[p2] {
					p--
				}
				p2 = p
			} else {
				p := p1 + 1
				for p < p2 && nums[p] == nums[p1] {
					p++
				}
				p1 = p
			}
		}
	}

	return res
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func choose(a, b, target int) int {
	if abs(a-target) > abs(b-target) {
		return b
	}
	return a
}
