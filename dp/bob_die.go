package dp

import (
	"fmt"
	"math"
)

// BobDie 返回 Bob 的生存率
func BobDie(n, m, i, j, k int) string {
	survive := BobDieRecur(n, m, i, j, k)
	return fmt.Sprintf("%d/%d", survive, int(math.Pow(4, float64(k))))
}

func BobDieRecur(n, m, i, j, k int) int {
	if i > n || i < 0 || j > m || j < 0 {
		return 0
	}
	if k == 0 {
		return 1
	}
	return BobDieRecur(n, m, i+1, j, k-1) +
		BobDieRecur(n, m, i, j+1, k-1) +
		BobDieRecur(n, m, i-1, j, k-1) +
		BobDieRecur(n, m, i, j-1, k-1)
}

// BobDieDP DP版本
func BobDieDP(n, m, i, j, k int) string {
	dp := make([][][]int, k+1)
	for h := 0; h < k+1; h++ {
		dp[h] = make([][]int, n+1)
		for x := 0; x < n+1; x++ {
			dp[h][x] = make([]int, m+1)
			for y := 0; y < m+1; y++ {
				if h == 0 {
					dp[h][x][y] = 1
				}
			}
		}
	}
	for h := 1; h < k+1; h++ {
		for x := 0; x < n+1; x++ {
			for y := 0; y < m+1; y++ {
				dp[h][x][y] += getValue(dp, n, m, x-1, y, h-1)
				dp[h][x][y] += getValue(dp, n, m, x+1, y, h-1)
				dp[h][x][y] += getValue(dp, n, m, x, y-1, h-1)
				dp[h][x][y] += getValue(dp, n, m, x, y+1, h-1)
			}
		}
	}
	return fmt.Sprintf("%d/%d", dp[k][i][j], int(math.Pow(4, float64(k))))
}

func getValue(dp [][][]int, n, m, i, j, k int) int {
	if i > n || i < 0 || j > m || j < 0 {
		return 0
	}
	return dp[k][i][j]
}
