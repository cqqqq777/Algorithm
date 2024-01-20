package day07

// LongestNoRepeat 最长无重复子串
func LongestNoRepeat(str string) int {
	if str == "" {
		return 0
	}
	// 记录上一次出现此字符的位置，使用数组替代 map
	lastIndex := make([]int, 255)
	for i := 0; i < len(lastIndex); i++ {
		lastIndex[i] = -1
	}
	pre, cur, res := -1, 0, 0
	for i := 0; i < len(str); i++ {
		pre = max(pre, lastIndex[str[i]])
		cur = i - pre
		res = max(cur, res)
		lastIndex[str[i]] = i
	}
	return res
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	lastIndex := make([]int, 255)
	for i := 0; i < 255; i++ {
		lastIndex[i] = -1
	}
	pre, result := -1, 0
	for k, v := range s {
		pre = max(pre, lastIndex[v])
		cur := k - pre
		lastIndex[v] = k
		result = max(result, cur)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
