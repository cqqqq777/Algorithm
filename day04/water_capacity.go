package day04

// WaterCapacity 给定容器直方图，返回容量
// 局部观察每一个位置能够容纳多少水
func WaterCapacity(arr []int) int {
	if arr == nil || len(arr) <= 2 {
		return 0
	}
	size := len(arr)
	leftMax, rightMax := make([]int, size), make([]int, size)
	for i := 1; i < size; i++ {
		leftMax[i] = max(leftMax[i-1], arr[i-1])
	}
	for i := size - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], arr[i+1])
	}
	res := 0
	for i := 0; i < size; i++ {
		res += max(min(leftMax[i], rightMax[i])-arr[i], 0)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
