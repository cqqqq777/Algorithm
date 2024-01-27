package winter09

// 峰值元素是指其值严格大于左右相邻值的元素。
// 给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)>>1 + l
		if nums[mid] > nums[mid+1] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
