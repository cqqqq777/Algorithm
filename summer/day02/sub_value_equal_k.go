package day02

// AllPair 返回数组中差值为 k 的去重数字对
// 利用 map，将 arr 所有数据放入 map 中，遍历 map 看与之对应差值为 k 的数是否在 map 中
func AllPair(arr []int, k int) [][]int {
	nums := make(map[int]struct{}, len(arr))
	for _, v := range arr {
		nums[v] = struct{}{}
	}
	res := make([][]int, 0)
	for key := range nums {
		if _, ok := nums[key+k]; ok {
			res = append(res, [][]int{{key, key + k}}...)
		}
	}
	return res
}
