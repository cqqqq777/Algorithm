package winter13

// 给定一个含有 n 个正整数的数组和一个正整数 target 。
// 找出该数组中满足其总和大于等于 target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	var sum int
	res := len(nums) + 1
	for r < len(nums) {
		sum += nums[r]
		for sum >= target && l < len(nums) {
			res = min(res, r-l+1)
			sum -= nums[l]
			l++
		}
		r++
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}
