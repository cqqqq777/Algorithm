package winter02

// 给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
// 请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
func merge(nums1 []int, m int, nums2 []int, n int) {
	if nums1 == nil || nums2 == nil {
		return
	}
	ptr1, ptr2, cur := m-1, n-1, m+n-1
	for ptr1 >= 0 && ptr2 >= 0 {
		if nums1[ptr1] > nums2[ptr2] {
			nums1[cur] = nums1[ptr1]
			ptr1--
		} else {
			nums1[cur] = nums2[ptr2]
			ptr2--
		}
		cur--
	}
	for ptr1 >= 0 {
		nums1[cur] = nums1[ptr1]
		ptr1--
		cur--
	}
	for ptr2 >= 0 {
		nums1[cur] = nums2[ptr2]
		ptr2--
		cur--
	}
}
