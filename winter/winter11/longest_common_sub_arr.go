package winter11

// 给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。
func findLength(nums1 []int, nums2 []int) int {
	// dp[i][j] 表示 nums1 以 i 位置结尾，nums2 以 j 位置结尾的最长子数组长度
	dp := make([][]int, len(nums1))
	dp[0] = make([]int, len(nums2))
	if nums1[0] == nums2[0] {
		dp[0][0] = 1
	}
	res := dp[0][0]
	for i := 1; i < len(nums2); i++ {
		if nums1[0] == nums2[i] {
			dp[0][i] = 1
		}
		res = max(dp[0][i], res)
	}
	for i := 1; i < len(nums1); i++ {
		dp[i] = make([]int, len(nums2))
		if nums1[i] == nums2[0] {
			dp[i][0] = 1
		}
		res = max(dp[i][0], res)
	}
	for i := 1; i < len(nums1); i++ {
		for j := 1; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			res = max(dp[i][j], res)
		}
	}
	return res
}
