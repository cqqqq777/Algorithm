package day01

// CoverMaxPoint 给定有序数组和长度，求覆盖最多的点个数
func CoverMaxPoint(arr []int, length int) int {
	return coverMaxPoint2(arr, length)
}

// 思路1：遍历数组，将 arr[i] 作为终点，看能覆盖前面多少个点，记录最大值
// 时间复杂度 O(N^2)
func coverMaxPoint1(arr []int, length int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		// 二分找到大于 arr[i] - length 最右边的点，更新 max
		l, r, min := 0, i, arr[i]-length
		for l < r {
			mid := l + (r-l)>>1
			if arr[mid] <= min {
				l = mid + 1
			} else {
				r = mid
			}
		}
		if max < i-l+1 {
			max = i - l + 1
		}
	}
	return max
}

// 思路2：利用窗口的思想
// 时间复杂度 O(N)
func coverMaxPoint2(arr []int, length int) int {
	max := 0
	for l := 0; l < len(arr); l++ {
		r := l
		for r < len(arr) && arr[r]-arr[l] <= length {
			r++
		}
		if max < r-l {
			max = r - l
		}
	}
	return max
}
