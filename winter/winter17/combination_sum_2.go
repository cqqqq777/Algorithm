package winter17

import (
	"sort"
)

// 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用 一次 。
// 注意：解集不能包含重复的组合。
func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	track := make([]int, 0)
	res := make([][]int, 0)
	var back func(idx, n int)
	back = func(idx, n int) {
		if n == target {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		if n > target {
			return
		}
		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			n += candidates[i]
			track = append(track, candidates[i])
			back(i+1, n)
			track = track[:len(track)-1]
			n -= candidates[i]
		}
	}
	back(0, 0)
	return res
}
