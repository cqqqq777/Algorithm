package winter07

import (
	"strconv"
	"strings"
)

// 给你两个版本号 version1 和 version2 ，请你比较它们。
//
// 版本号由一个或多个修订号组成，各修订号由一个 '.' 连接。每个修订号由 多位数字 组成，可能包含 前导零 。每个版本号至少包含一个字符。
// 修订号从左到右编号，下标从 0 开始，最左边的修订号下标为 0 ，下一个修订号下标为 1 ，以此类推。例如，2.5.33 和 0.1 都是有效的版本号。
//
// 比较版本号时，请按从左到右的顺序依次比较它们的修订号。比较修订号时，只需比较 忽略任何前导零后的整数值 。
// 也就是说，修订号 1 和修订号 001 相等 。如果版本号没有指定某个下标处的修订号，则该修订号视为 0 。
// 例如，版本 1.0 小于版本 1.1 ，因为它们下标为 0 的修订号相同，而下标为 1 的修订号分别为 0 和 1 ，0 < 1 。
// 如果 version1 > version2 返回 1，
// 如果 version1 < version2 返回 -1，
// 除此之外返回 0。
func compareVersion(version1 string, version2 string) int {
	nums1, nums2 := strings.Split(version1, "."), strings.Split(version2, ".")
	var m, n int
	for m < len(nums1) && n < len(nums2) {
		num1, _ := strconv.Atoi(nums1[m])
		num2, _ := strconv.Atoi(nums2[n])
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		} else {
			m++
			n++
		}
	}
	if len(nums1) > len(nums2) {
		for ; m < len(nums1); m++ {
			if num, _ := strconv.Atoi(nums1[m]); num != 0 {
				return 1
			}
		}
	}
	if len(nums2) > len(nums1) {
		for ; n < len(nums2); n++ {
			if num, _ := strconv.Atoi(nums2[n]); num != 0 {
				return -1
			}
		}
	}
	return 0
}
