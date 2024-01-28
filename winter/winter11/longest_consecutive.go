package winter11

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
func longestConsecutive(nums []int) int {
	numSet := make(map[int]struct{})
	for _, v := range nums {
		numSet[v] = struct{}{}
	}
	var res int
	for v := range numSet {
		if _, ok := numSet[v-1]; !ok {
			curRes := 1
			next := v + 1
			_, ok1 := numSet[next]
			for ; ok1; curRes = curRes + 1 {
				next++
				_, ok1 = numSet[next]
			}
			res = max(res, curRes)
		}
	}
	return res
}
