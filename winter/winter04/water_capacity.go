package winter04

// 给定直方图，求能接多少雨水
func trap(height []int) int {
	if height == nil || len(height) <= 2 {
		return 0
	}
	size := len(height)
	leftMax, rightMax := make([]int, size), make([]int, size)
	for i := 1; i < size; i++ {
		leftMax[i] = max(leftMax[i-1], height[i-1])
	}
	for i := size - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i+1])
	}
	res := 0
	for i := 0; i < size; i++ {
		res += max(min(leftMax[i], rightMax[i])-height[i], 0)
	}
	return res
}
