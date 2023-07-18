package recursion

func WinnerScore(arr []int) int {
	if arr == nil {
		return 0
	}
	f := first(arr, 0, len(arr)-1)
	s := second(arr, 0, len(arr)-1)
	if f > s {
		return f
	}
	return s
}

// first 表示先手拿牌
func first(arr []int, l, r int) int {
	if l == r {
		return arr[l]
	}
	// 已经先手拿牌了，接下来应该由后手拿牌的人先拿了再拿
	left := arr[l] + second(arr, l+1, r)  // 拿了左边的牌
	right := arr[r] + second(arr, l, r-1) // 拿了右边的牌
	if left > right {
		return left
	}
	return right
}

// second 表示后手拿牌
func second(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	left := first(arr, l+1, r)  // 表示先手的人把 l 位置的牌拿了
	right := first(arr, l, r-1) // 表示先手的人把 r 位置的牌拿了
	// 因为先手的人肯定把大牌取走了，所以这里返回较小值
	if left < right {
		return left
	}
	return right
}
