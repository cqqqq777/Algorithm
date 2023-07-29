package recursion

// WalkMethods 返回所有方法数
func WalkMethods(N, M, P, K int) int {
	return f(N, M, P, K)
}

// cur 当前位置
// rest 还剩多少步可以走
// E 结束位置
func f(N, cur, E, rest int) int {
	if rest == 0 {
		if cur == E {
			return 1
		}
		return 0
	}
	if cur == 1 {
		return f(N, 2, E, rest-1)
	}
	if cur == N {
		return f(N, N-1, E, rest-1)
	}
	return f(N, cur-1, E, rest-1) + f(N, cur+1, E, rest-1)
}

// WalkMethods2 对机器人走路的优化
func WalkMethods2(N, M, P, K int) int {
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, K+1)
		for j := 0; i < K+1; j++ {
			dp[i][j] = -1
		}
	}
	return f2(N, M, P, K, dp)
}

func f2(N, cur, E, rest int, dp [][]int) int {
	// 命中缓存，直接返回
	if dp[cur][rest] != -1 {
		return dp[cur][rest]
	}
	// 缓存未命中，每次计算结果计入缓存
	if rest == 0 {
		if cur == E {
			dp[cur][rest] = 1
		} else {
			dp[cur][rest] = 0
		}
		return dp[cur][rest]
	}
	if cur == 1 {
		dp[cur][rest] = f(N, 2, E, rest-1)
	} else if cur == N {
		dp[cur][rest] = f(N, N-1, E, rest-1)
	} else {
		dp[cur][rest] = f(N, cur-1, E, rest-1) + f(N, cur+1, E, rest-1)
	}

	return dp[cur][rest]
}
