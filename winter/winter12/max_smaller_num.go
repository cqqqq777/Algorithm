package winter12

import (
	"sort"
)

// 数组A中给定可以使用的1~9的数，返回由A数组中的元素组成的小于n（n > 0）的最大数。
// 例如：A = {1, 2, 4, 9}，n = 2533，返回2499。
func maxSmaller(nums []int, n int) int {
	if len(nums) == 0 {
		return -1
	}

	// 排序后方便二分查找
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	if n <= nums[0] {
		return -1
	}

	// 把 n 转化为数组后面好比较
	target := make([]int, 0)
	for n > 0 {
		target = append(target, n%10)
		n /= 10
	}
	for i := 0; i < len(target)/2; i++ {
		target[i], target[len(target)-1-i] = target[len(target)-1-i], target[i]
	}

	// track 记录选择的路径
	track := make([]int, 0)
	var back func(idx int) bool
	back = func(idx int) bool {
		if idx == len(target) {
			return true
		}

		if target[idx] < nums[0] {
			return false
		}

		// 每次选择小于等于当前位的最大值，加在路径里面
		cur := bs(nums, target[idx])
		// 来到最后一位，不能相等
		if idx == len(target)-1 && nums[cur] == target[idx] {
			// cur 已经是 nums 里面最小的了，所以这条路走不通
			if cur == 0 {
				return false
			}
			// 还有比 cur 更小的，加在路径里面返回即可
			track = append(track, nums[cur-1])
			return true
		}
		track = append(track, nums[cur])
		// cur 比当前位要小，直接把 nums 里面最大的数全部往后面加即可
		if nums[cur] < target[idx] {
			for i := idx + 1; i < len(target); i++ {
				track = append(track, nums[len(nums)-1])
			}
			return true
		}
		// 进入下一位
		// 如果选出结果了，直接往上返回即可
		if back(idx + 1) {
			return true
		}
		// 没选出结果，开始回溯
		track = track[:len(track)-1]
		// cur 已经是 nums 里面最小的数了，往上回溯
		if cur == 0 {
			return false
		}
		// 还有比 cur 更小的数，加在路径里面即可
		track = append(track, nums[cur-1])
		for i := idx + 1; i < len(target); i++ {
			track = append(track, nums[len(nums)-1])
		}
		return true
	}
	back(0)
	var res int
	// 将 track 转化为结果
	// 不能在相同的位数内选出结果，那答案就是少一位，全选 nums 里面最大的数
	if len(track) == 0 {
		for i := 0; i < len(target)-1; i++ {
			track = append(track, nums[len(nums)-1])
		}
	}
	for i := 0; i < len(track); i++ {
		res = res*10 + track[i]
	}
	return res
}

// 二分查找小于等于 target 的最大数
func bs(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)>>1 + l
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}
