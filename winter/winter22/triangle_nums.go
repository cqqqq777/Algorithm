package winter22

import "sort"

// 给定一个包含非负整数的数组 nums ，返回其中可以组成三角形三条边的三元组个数。
func triangleNumber(nums []int) (ans int) {
	if len(nums) <= 2 {
		return 0
	}

	sort.Ints(nums)

	var res int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == 0 {
			continue
		}
		for p1 := i + 1; p1 < len(nums)-1; p1++ {
			p2 := p1 + 1
			for p2 < len(nums) && nums[p1]+nums[i] > nums[p2] {
				p2++
			}
			res += p2 - p1 - 1
		}
	}
	return res
}
