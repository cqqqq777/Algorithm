package winter07

import "math"

// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
func subsets(nums []int) [][]int {
	// 幂集有 2^n 个
	res := make([][]int, 0, int(math.Pow(2, float64(len(nums)))))

	track := make([]int, 0)
	var back func(set []int, idx int)
	back = func(set []int, idx int) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)
		for i := idx; i < len(nums); i++ {
			track = append(track, nums[i])
			back(track, i+1)
			track = track[:len(track)-1]
		}
	}
	back(track, 0)
	return res
}
