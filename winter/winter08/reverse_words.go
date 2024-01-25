package winter08

import "strings"

// 给你一个字符串 s ，请你反转字符串中 单词 的顺序。
// 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
// 返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
// 注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
func reverseWords(s string) string {
	res := make([]string, 0)
	// 双指针
	var p1, p2 int
	// pre 记录当前位置前一个字符
	pre := s[0]
	for i := 1; i < len(s); i++ {
		p2++
		if s[i] == ' ' {
			// 如果当前位置是空格并且前一个位置不是空格
			// 将双指针内的字符串加入结果中
			if pre != ' ' {
				res = append(res, s[p1:p2])
			}
		} else {
			// 当前字符不是空格且前一个字符是空格
			// 更新 p1
			if pre == ' ' {
				p1 = p2
			}
		}
		pre = s[i]
	}
	// 避免最后一个字符是空格的情况
	if s[len(s)-1] != ' ' {
		res = append(res, s[p1:p2+1])
	}

	// 反转 res
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return strings.Join(res, " ")
}
