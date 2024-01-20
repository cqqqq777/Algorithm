package day04

import (
	"container/heap"
)

type coffeeMachine struct {
	readyTime, workTime int
}

type priorityQueue []*coffeeMachine

func (p priorityQueue) Less(i, j int) bool {
	return p[i].workTime+p[i].readyTime < p[j].workTime+p[j].readyTime
}

func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p priorityQueue) Len() int {
	return len(p)
}

func (p *priorityQueue) Push(x interface{}) {
	cm := x.(*coffeeMachine)
	*p = append(*p, cm)
}

func (p *priorityQueue) Pop() interface{} {
	old := *p
	cm := old[len(old)-1]
	*p = old[:len(old)-1]
	return cm
}

// WashCupMinTime
// 数组arr：表示几个咖啡机，这几个咖啡机生产一杯咖啡所需要的时间就是数组中的值，例如arr=[2,3,7]就表示第一台咖啡机生产一杯咖啡需要2单位时间，第二台需要3单位时间，第三台需要7单位时间。
// int N：表示有 N 个人需要用咖啡机制作咖啡，每人一杯，同时，假设制作完咖啡后，喝咖啡时间为0，一口闷。
// int a：表示用洗碗机洗一个咖啡杯需要的时间，串行运行。
// int b：表示咖啡杯也可以不洗，自然晾干的时间。咖啡杯要么洗，要么晾干。
// 现在，请你求出这 N 个人从开始用咖啡杯制作咖啡到杯子洗好或者晾干的最少时间？
func WashCupMinTime(arr []int, N, a, b int) int {
	pq := make(priorityQueue, len(arr))
	for i := 0; i < len(arr); i++ {
		pq[i] = &coffeeMachine{workTime: arr[i]}
	}
	heap.Init(&pq)
	drinks := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		cm := heap.Pop(&pq).(*coffeeMachine)
		drinks[i] = cm.workTime + cm.readyTime
		cm.readyTime += cm.workTime
		heap.Push(&pq, cm)
	}
	return washTime(drinks, a, b, 0, 0)
}

// 返回洗完 drinks[index..N-1] 杯子的最少时间
// washLine: 洗杯子的机器在 washLine 时间为空闲
func washTime(drinks []int, a, b, index, washLine int) int {
	if index == len(drinks)-1 {
		return min(max(washLine, drinks[index])+a, drinks[index]+b)
	}
	// 选择使用机器洗杯子
	wash := max(washLine, drinks[index]) + a
	time1 := washTime(drinks, a, b, index+1, wash)
	// 选择自然挥发
	dry := drinks[index] + b
	time2 := washTime(drinks, a, b, index+1, washLine)
	return min(max(wash, time1), max(dry, time2))
}

func WashCupMinTime2(arr []int, N, a, b int) int {
	pq := make(priorityQueue, len(arr))
	for i := 0; i < len(arr); i++ {
		pq[i] = &coffeeMachine{workTime: arr[i]}
	}
	heap.Init(&pq)
	drinks := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		cm := heap.Pop(&pq).(*coffeeMachine)
		drinks[i] = cm.workTime + cm.readyTime
		cm.readyTime += cm.workTime
		heap.Push(&pq, cm)
	}
	return washTimeDP(drinks, a, b)
}

// 动态规划
// 在上面的递归中 index 和 washLine 一确定，最少时间也确定
func washTimeDP(drinks []int, a, b int) int {
	if drinks == nil || len(drinks) < 1 {
		return 0
	}
	size := len(drinks)
	maxWashLine := drinks[size-1] + a*size // 给 washLine 分配一个足够大的值
	dp := make([][]int, size)
	dp[size-1] = make([]int, maxWashLine)
	for i := 0; i < maxWashLine; i++ {
		dp[size-1][i] = min(max(i, drinks[size-1])+a, drinks[size-1]+b)
	}
	for i := size - 2; i >= 0; i-- {
		maxWashLine = drinks[i] + (i+1)*a
		dp[i] = make([]int, maxWashLine)
		for j := 0; j < maxWashLine; j++ {
			dp[i][j] = min(max(j, drinks[i])+a, drinks[i]+b)
		}
	}
	return dp[0][0]
}
