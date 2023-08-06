package day05

import "sort"

type job struct {
	money, hard int // 报酬与难度
}

// ChooseWork
// 为了找到自己满意的工作，牛牛收集了每种工作的难度和报酬。
// 牛牛选工作的标准是在难度不超过自身能力值的情况下，牛牛选择报酬最高的工作。
// 在牛牛选定了自己的工作后，牛牛的小伙伴们来找牛牛帮忙选工作，牛牛依然使用自己的标准来帮助小伙伴们。牛牛的小伙伴太多了，于是他只好把这个任务交给了你
// 给定一个Job类型的数组 jobArr，表示所有的工作。给定一个 int 类型的数组 arr，表示所有小伙伴的能力。
// 返回 int 类型的数组，表示每一个小伙伴按照牛牛的标准选工作后所能获得的报酬。
// 每个工作可以被选择多次
// TODO: use sorted map to decrease time complexity
func ChooseWork(jobArr []*job, ability []int) []int {
	// 将每个工作先按难度升序排序，如果难度相同就将其按报酬降序排序
	sort.Slice(jobArr, func(i, j int) bool {
		if jobArr[i].hard == jobArr[j].hard {
			return jobArr[i].money < jobArr[j].money
		}
		return jobArr[i].hard < jobArr[j].hard
	})
	// 只保留每个难度报酬最高的工作
	// 去掉难度增加但报酬未增加的工作
	jobMap := make(map[int]int, 0)
	jobMap[jobArr[0].hard] = jobArr[0].money
	pre := jobArr[0]
	for i := 1; i < len(jobArr); i++ {
		if pre.hard != jobArr[i].hard && jobArr[i].money > pre.money {
			jobMap[jobArr[i].hard] = jobArr[i].money
		}
	}
	jobs := make([]int, len(jobMap))
	i := 0
	for k := range jobMap {
		jobs[i] = k
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i] < jobs[j]
	})
	// 每个人都选择刚好低于自身能力的工作
	res := make([]int, len(ability))
	for i = 0; i < len(ability); i++ {
		if ability[i] < jobs[0] {
			continue
		}
		for j := 1; j < len(jobs); j++ {
			if jobs[j] > ability[j] {
				res[i] = jobMap[jobs[j-1]]
				break
			}
		}
	}
	return res
}
