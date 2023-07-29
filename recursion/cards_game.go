package recursion

func WinnerScore(arr []int) int {
	if arr == nil {
		return 0
	}
	f := first(arr, 0, len(arr)-1)
	s := second(arr, 0, len(arr)-1)
	if f > s {
		return f
	}
	return s
}

// first 表示先手拿牌
func first(arr []int, l, r int) int {
	if l == r {
		return arr[l]
	}
	// 已经先手拿牌了，接下来应该由后手拿牌的人先拿了再拿
	left := arr[l] + second(arr, l+1, r)  // 拿了左边的牌
	right := arr[r] + second(arr, l, r-1) // 拿了右边的牌
	if left > right {
		return left
	}
	return right
}

// second 表示后手拿牌
func second(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	left := first(arr, l+1, r)  // 表示先手的人把 l 位置的牌拿了
	right := first(arr, l, r-1) // 表示先手的人把 r 位置的牌拿了
	// 因为先手的人肯定把大牌取走了，所以这里返回较小值
	if left < right {
		return left
	}
	return right
}

// WinnerScoreDP 表示动态规划解法
func WinnerScoreDP(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	l, r := len(arr), len(arr)
	f, s := make([][]int, l), make([][]int, r)
	for i := 0; i < len(arr); i++ {
		f[i], s[i] = make([]int, len(arr)), make([]int, len(arr))
	}
	for i := l - 1; i >= 0; i-- {
		for j := i; j < r; j++ {
			if i == j {
				f[i][j] = arr[i]
				s[i][j] = 0
			} else {
				f[i][j] = max(arr[i]+s[i+1][j], arr[j]+s[i][j-1])
				s[i][j] = min(f[i+1][j], f[i][j-1])
			}
		}
	}
	return max(f[0][len(arr)-1], s[0][len(arr)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
