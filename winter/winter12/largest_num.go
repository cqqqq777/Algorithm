package winter12

import (
	"sort"
	"strconv"
	"strings"
)

// 给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
// 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		// x 拼接到 y 前面要乘 fx
		fx, fy := 1, 1
		for y > 0 {
			fx *= 10
			y /= 10
		}
		for x > 0 {
			fy *= 10
			x /= 10
		}
		return nums[i]*fx+nums[j] > nums[j]*fy+nums[i]
	})
	var res strings.Builder
	for _, v := range nums {
		res.WriteString(strconv.Itoa(v))
	}
	return res.String()
}
