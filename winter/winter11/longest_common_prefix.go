package winter11

// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
func longestCommonPrefix(strs []string) string {
	// 分治
	// 每次找两个字符串的公共前缀，再对公共前缀找公共前缀
	if len(strs) == 1 {
		return strs[0]
	}
	if len(strs) == 2 {
		return twoStrLongestCommonPrefix(strs[0], strs[1])
	}
	l, r := 0, len(strs)-1
	mid := (r-l)>>1 + l
	pre1, pre2 := longestCommonPrefix(strs[l:mid]), longestCommonPrefix(strs[mid:r+1])
	return twoStrLongestCommonPrefix(pre1, pre2)
}

// 返回两个字符串的最长公共前缀
func twoStrLongestCommonPrefix(str1, str2 string) string {
	p1, p2 := 0, 0
	for ; p1 < len(str1) && p2 < len(str2); p1, p2 = p1+1, p2+1 {
		if str1[p1] != str2[p2] {
			break
		}
	}
	return str1[0:p1]
}
