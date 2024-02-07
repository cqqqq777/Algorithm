package shuffle

import "math/rand"

// 给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。打乱后，数组的所有排列应该是 等可能 的。
// 实现 Solution class:
// Solution(int[] nums) 使用整数数组 nums 初始化对象
// int[] reset() 重设数组到它的初始状态并返回
// int[] shuffle() 返回数组随机打乱后的结果

type Solution struct {
	nums, origin []int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums:   nums,
		origin: append([]int(nil), nums...),
	}
}

func (this *Solution) Reset() []int {
	copy(this.nums, this.origin)
	return this.nums
}

func (this *Solution) Shuffle() []int {
	for i := len(this.nums) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		this.nums[i], this.nums[r] = this.nums[r], this.nums[i]
	}
	return this.nums
}
