package winter02

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
func permute(nums []int) [][]int {
	if nums == nil || len(nums) == 1 {
		return [][]int{nums}
	}
	tmp, res := make([]int, 0), make([][]int, 0)
	used := make(map[int]bool)

	var permuteRecur func(int)

	permuteRecur = func(index int) {
		if index == len(nums) {
			c := make([]int, len(tmp))
			copy(c, tmp)
			res = append(res, c)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			tmp = append(tmp, nums[i])
			used[i] = true
			permuteRecur(index + 1)
			tmp = tmp[:len(tmp)-1]
			used[i] = false
		}
	}
	permuteRecur(0)
	return res
}
