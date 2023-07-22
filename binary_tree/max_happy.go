package binary_tree

import "math"

type Employee struct {
	happy        int         // 这名员工可以带来的快乐值
	subordinates []*Employee // 这名员工有哪些直接下级
}

type HappyInfo struct {
	YHappy, NHappy int
}

// MaxHappy 返回最大快乐值
func MaxHappy(boss *Employee) int {
	if boss == nil {
		return 0
	}
	info := MaxHappyRecur(boss)
	return int(math.Max(float64(info.YHappy), float64(info.NHappy)))
}

func MaxHappyRecur(node *Employee) *HappyInfo {
	info := new(HappyInfo)
	if node.subordinates == nil {
		info.YHappy = node.happy
		info.NHappy = 0
		return info
	}
	info.YHappy, info.NHappy = node.happy, 0
	for _, v := range node.subordinates {
		nextInfo := MaxHappyRecur(v)
		info.YHappy += nextInfo.NHappy
		info.NHappy += int(math.Max(float64(nextInfo.YHappy), float64(nextInfo.NHappy)))
	}
	return info
}
