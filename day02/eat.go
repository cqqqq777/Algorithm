package day02

// Winner 先手胜利返回1，后手胜利返回0
// 每个人都必须吃4的 x 次幂数量，不能吃到有效数量的玩家失败
// 每个人都是按照最佳方法进行游戏
func Winner(n int) int {
	vis := make(map[int]int)
	return winner(n, vis)
}

func w(n int) int {
	if n < 5 {
		if n == 0 || n == 2 {
			return 0
		}
		return 1
	}
	base := 1
	for base <= n {
		if w(n-base) == 0 {
			return 1
		}
		// 防溢出
		if base > n/4 {
			break
		}
		base *= 4
	}
	return 0
}

// 记忆化搜索
func winner(n int, vis map[int]int) int {
	if v, ok := vis[n]; ok {
		return v
	}
	if n < 5 {
		if n == 0 || n == 2 {
			vis[n] = 0
		} else {
			vis[n] = 1
		}
		return vis[n]
	}
	base := 1
	for base <= n {
		if winner(n-base, vis) == 0 {
			vis[n-base] = 0
			vis[n] = 1
			return 1
		}
		// 防溢出
		if base > n/4 {
			break
		}
		base *= 4
	}
	vis[n] = 0
	return 0
}

// 根据上面的方法输出结果打表发现规律
func table(n int) int {
	if n%5 == 0 || n%5 == 2 {
		return 0
	}
	return 1
}
