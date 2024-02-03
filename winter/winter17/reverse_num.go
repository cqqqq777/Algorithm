package winter17

import (
	"math"
	"strconv"
)

// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
// 假设环境不允许存储 64 位整数（有符号或无符号）。
func reverse(x int) int {
	if x == 0 {
		return 0
	}
	var flag bool
	if x < 0 {
		flag = true
		x = -x
	}
	// 负数比正数多一个，所以全部转化成负数考虑
	data := strconv.Itoa(x)
	var res int
	for i := len(data) - 1; i >= 0; i-- {
		cur := int(data[i] - '0')
		if math.MinInt32-res*10 > cur {
			return 0
		}
		res = res*10 - cur
	}
	if flag {
		return res
	}
	if res == math.MinInt32 {
		return 0
	}
	return -res
}
