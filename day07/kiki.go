package day07

import "math"

// Kiki
// CC里面有一个土豪很喜欢一位女直播Kiki唱歌，平时就经常给她点赞、送礼、私聊。最近CC直播平台在举行中秋之星主播唱歌比赛，假设一开始该女主播的初始人气值为start，能够晋升下一轮人气需要刚好达到end，土豪给主播增加人气的可以采取的方法有：
// a. 点赞 花费x C币，人气 + 2
// b. 送礼 花费y C币，人气 * 2
// c. 私聊 花费z C币，人气 - 2
// 其中 end 远大于start且end为偶数， 请写一个程序帮助土豪计算一下，最少花费多少C币就能帮助该主播Kiki将人气刚好达到end，从而能够晋级下一轮？
// 输入描述：
// 第一行输入5个数据，分别为：x y z start end，每项数据以空格分开。其中：0＜x, y, z＜＝10000， 0＜start, end＜＝1000000
// 输出描述：
// 需要花费的最少C币。
// 示例1:
// 输入
// 3 100 1 2 6
// 输出
// 6
func Kiki(start, end, add, times, del int) int {
	if start == end {
		return 0
	}
	return minCoin(0, end, add, times, del, 2*end, ((end-start)/2)*add)
}

func minCoin(pre, end, add, times, del, limitEnd, limitCoin int) int {
	if pre > limitCoin || end < 0 || end > limitEnd {
		return math.MaxInt
	}
	if pre == end {
		return pre
	}
	m := math.MaxInt
	p1 := minCoin(pre+add, end-2, add, times, del, limitEnd, limitCoin)
	if p1 != math.MaxInt {
		m = p1
	}
	p2 := minCoin(pre+times, end/2, add, times, del, limitEnd, limitCoin)
	if p2 != math.MaxInt {
		m = min(m, p2)
	}
	p3 := minCoin(pre+del, end+2, add, times, del, limitEnd, limitCoin)
	if p3 != math.MaxInt {
		m = min(m, p3)
	}
	return m
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minCoinDP(start, end, add, times, del int) int {
	if start == end {
		return 0
	}
	limitEnd, limitCoin := end*2, ((end-start)/2)*add
	dp := make([][]int, limitCoin+1)
	for pre := 0; pre <= limitCoin; pre++ {
		dp[pre] = make([]int, limitEnd+1)
		for e := 0; e < limitEnd; e++ {
			if e == start {
				dp[pre][e] = pre
			} else {
				dp[pre][e] = math.MaxInt
			}
		}
	}
	for pre := limitCoin; pre >= 0; pre-- {
		for e := 0; e < limitEnd; e++ {
			if e-2 >= 0 && pre+add <= limitCoin {
				dp[pre][e] = min(dp[pre][e], dp[pre+add][e-2])
			}
			if e+2 < limitEnd && pre+del <= limitCoin {
				dp[pre][e] = min(dp[pre][e], dp[pre+del][e+2])
			}
			if e&1 == 0 {
				if e/2 >= 0 && pre+times <= limitCoin {
					dp[pre][e] = min(dp[pre][e], dp[pre+times][e/2])
				}
			}
		}
	}
	return dp[0][end]
}
