package winter15

import "sort"

// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
func permuteUnique(nums []int) [][]int {
	// 对 nums 排序，让相同元素的相对次序保持不变
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	used, track := make(map[int]bool), make([]int, 0)
	res := make([][]int, 0)
	var back func(idx int)
	back = func(idx int) {
		if idx == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			// 保证相同元素的相对次序不变
			// 如果当前元素和前一个元素相等，则必须要求前一个元素被选过了才能选当前元素
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			track = append(track, nums[i])
			back(idx + 1)
			track = track[:len(track)-1]
			used[i] = false
		}
	}
	back(0)
	return res
}
