package winter19

// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
// 每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
// 0 <= j <= nums[i]
// i + j < n
// 返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	// 每次都尽量用更少的次数来到更远的地方
	res, end, maxPos := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		// maxPos 表示前面已经遍历过的点能够到达的最远的地方
		maxPos = max(maxPos, i+nums[i])
		// i 来到了前面的点能够到达的最远的地方，就跳到这里来，结果加一
		if i == end {
			res++
			end = maxPos
		}
	}
	return res
}
