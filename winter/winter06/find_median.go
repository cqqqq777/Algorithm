package winter06

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
// 划分区间法
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	iMin, iMax := 0, m
	for iMin <= iMax {
		// i 和 j 分别表示两个数组的右区间的最小索引
		i := (iMax-iMin)>>1 + iMin
		j := (m+n+1)>>1 - i
		if j > 0 && i < m && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i > 0 && j < n && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			var maxLeft, minRight int
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}

			if (m+n)&1 == 1 {
				return float64(maxLeft)
			}

			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = min(nums1[i], nums2[j])
			}
			return float64((maxLeft + minRight) / 2)
		}
	}
	return 0
}
