package winter13

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
// 子数组是数组中元素的连续非空序列。
func subArraySum(nums []int, k int) int {
	var res, pre int
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		res += m[pre-k]
		m[pre]++
	}
	return res
}
