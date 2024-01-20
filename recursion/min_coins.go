package recursion

import "math"

// MinCoins 换钱最少货币数
func MinCoins(arr []int, aim int) int {
	if arr == nil {
		return -1
	}
	if aim == 0 {
		return 0
	}
	for i := 1; i <= aim; i++ {
		if MinCoinsF(aim, i, arr) {
			return i
		}
	}
	return -1
}

func MinCoinsF(aim, rest int, arr []int) bool {
	if rest == 0 {
		if aim == 0 {
			return true
		}
		return false
	}
	for i := 0; i < len(arr); i++ {
		if MinCoinsF(aim-arr[i], rest-1, arr) {
			return true
		}
	}
	return false
}

func MinCoins2(arr []int, aim int) int {
	if arr == nil {
		return -1
	}
	if aim == 0 {
		return 0
	}
	dp := make([]map[int]bool, aim+1)
	dp[0] = make(map[int]bool)
	for i := 1; i <= aim; i++ {
		dp[i] = make(map[int]bool)
		if MinCoinsF2(aim, i, arr, dp) {
			return i
		}
	}
	return -1
}

func MinCoinsF2(aim, rest int, arr []int, dp []map[int]bool) bool {
	v, ok := dp[rest][aim]
	if ok {
		return v
	}
	if rest == 0 {
		if aim == 0 {
			dp[rest][aim] = true
		} else {
			dp[rest][aim] = false
		}
		return dp[rest][aim]
	}
	for i := 0; i < len(arr); i++ {
		if MinCoinsF2(aim-arr[i], rest-1, arr, dp) {
			dp[rest][aim] = true
			break
		}
	}
	if _, ok = dp[rest][aim]; !ok {
		dp[rest][aim] = false
	}
	return dp[rest][aim]
}

func MinCoins3(arr []int, aim int) int {
	if arr == nil || aim < 0 {
		return -1
	}
	dp := make([]int, aim+1)
	const INF = math.MaxInt32
	for i := 1; i <= aim; i++ {
		dp[i] = INF
	}
	for i := 1; i <= aim; i++ {
		for j := 0; j < len(arr); j++ {
			if i >= arr[j] && dp[i-arr[j]] != INF {
				dp[i] = min(dp[i], dp[i-arr[j]]+1)
			}
		}
	}
	if dp[aim] == INF {
		return -1
	}
	return dp[aim]
}
