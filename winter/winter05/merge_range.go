package winter05

import "sort"

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 先按左端点升序排序，使用 merged 数组存储结果，每次查看 merged 最后一个区间跟 intervals 中能否合并
	merged, i := make([][]int, 0, len(intervals)), 0
	merged = append(merged, intervals[0])
	for j := 1; j < len(intervals); j++ {
		if intervals[j][0] > merged[i][1] {
			merged = append(merged, intervals[j])
			i++
		} else {
			merged[i][1] = max(merged[i][1], intervals[j][1])
		}
	}
	return merged
}
