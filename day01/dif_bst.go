package day01

// DifBST 给定节点个数返回能形成的不同二叉树个数
func DifBST(nodes int) int {
	m := make(map[int]int)
	return process2(nodes, m)
}

// 暴力递归
func process1(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if n == 1 || n == 2 {
		return n
	}
	res := 0
	for i := 0; i < n; i++ {
		res += process1(n-i-1) * process1(i)
	}
	return res
}

// 记忆化搜索
func process2(n int, nums map[int]int) int {
	v, ok := nums[n]
	if ok {
		return v
	}
	if n < 0 {
		nums[n] = 0
	}
	if n == 0 || n == 1 {
		nums[n] = 1
	}
	if n == 2 {
		nums[n] = 2
	}
	if v, ok = nums[n]; ok {
		return v
	}
	res := 0
	for i := 0; i < n; i++ {
		leftNums, ok := nums[i]
		if !ok {
			leftNums = process2(i, nums)
			nums[i] = leftNums
		}
		rightNums, ok := nums[n-i-1]
		if !ok {
			rightNums = process2(n-i-1, nums)
			nums[n-i-1] = rightNums
		}
		res += leftNums * rightNums
	}
	nums[n] = res
	return res
}

// dp 版本
func process3(n int) int {
	if n < 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}
