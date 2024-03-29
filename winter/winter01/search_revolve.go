package winter01

// 整数数组 nums 按升序排列，数组中的值 互不相同 。
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，
// 使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
// 例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	// 旋转后的数组是局部有序的
	// 使用二分法每次确定一个局部有序的区间
	l, r := 0, len(nums)-1
	// 代表左半部分有序
	for l <= r {
		mid := (r-l)>>1 + l
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[0] {
			// 左半部分有序
			if nums[mid] > target && nums[0] <= target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// 右半部分有序
			if nums[mid] < target && nums[len(nums)-1] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
