package dp

// WinnerScoreDP 解决棋牌博弈问题
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
