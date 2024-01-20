package winter01

import "sort"

// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
// 你返回所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
// 排序 + 双指针
func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return [][]int{}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := make([][]int, 0)
	for k, v := range nums {
		// 因为已经排序，一旦出现 v > 0，后面不可能还有答案
		if v > 0 {
			return res
		}
		// 排除掉重复值
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		l, r := k+1, len(nums)-1
		for l < r {
			if v+nums[l]+nums[r] == 0 {
				res = append(res, []int{v, nums[l], nums[r]})
				// 排除重复值
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if v+nums[l]+nums[r] < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
