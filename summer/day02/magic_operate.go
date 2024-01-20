package day02

import "sort"

// MagicOpNums 返回一共能进行几次 magic 操作
func MagicOpNums(set1, set2 []int) int {
	sum1, sum2 := 0, 0
	for _, v := range set1 {
		sum1 += v
	}
	for _, v := range set2 {
		sum2 += v
	}
	if avg(sum1, len(set1)) == avg(sum2, len(set2)) {
		return 0
	}
	var moreSet, lessSet []int
	var moreSize, moreSum, lessSize, lessSum int
	if avg(sum1, len(set1)) > avg(sum2, len(set2)) {
		moreSet, moreSize, moreSum = set1, len(set1), sum1
		lessSet, lessSize, lessSum = set2, len(set2), sum2
	} else {
		moreSet, moreSize, moreSum = set2, len(set2), sum2
		lessSet, lessSize, lessSum = set1, len(set1), sum1
	}
	sort.Slice(moreSet, func(i, j int) bool {
		return moreSet[i] < moreSet[j]
	})
	mapLess := make(map[int]struct{}, lessSize)
	for _, v := range lessSet {
		mapLess[v] = struct{}{}
	}
	res := 0
	for i := 0; i < len(moreSet); i++ {
		cur := float64(moreSet[i])
		if cur > avg(lessSum, lessSize) && cur < avg(moreSum, moreSize) {
			if _, ok := mapLess[moreSet[i]]; !ok {
				moreSum -= moreSet[i]
				moreSize--
				lessSize++
				lessSum += moreSet[i]
				mapLess[moreSet[i]] = struct{}{}
				res++
			}
		}
	}
	return res
}

func avg(a, b int) float64 {
	return float64(a) / float64(b)
}
