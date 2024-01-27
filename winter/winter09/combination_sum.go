package winter09

// 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。
// 你可以按 任意顺序 返回这些组合。
// candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
// 对于给定的输入，保证和为 target 的不同组合数少于 150 个。
func combinationSum(candidates []int, target int) [][]int {
	res, track := make([][]int, 0), make([]int, 0)
	var back func(trackSum, idx int)
	back = func(trackSum, idx int) {
		if trackSum == target {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		if trackSum > target {
			return
		}
		for i := idx; i < len(candidates); i++ {
			track = append(track, candidates[i])
			back(trackSum+candidates[i], i)
			track = track[:len(track)-1]
		}
	}
	back(0, 0)
	return res
}
