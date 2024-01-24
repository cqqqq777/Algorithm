package winter07

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
// 如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	// 记录 t 中每个字符的个数
	tCount := [128]int{}
	for i := 0; i < len(t); i++ {
		tCount[t[i]]++
	}

	// 记录滑动窗口中的字符数和与 t 的差异数
	windowCount, diff := [128]int{}, len(t)

	left, minIdx, minLen := 0, -1, len(s)+1
	for right := 0; right < len(s); right++ {
		// t 中不包含此字符，不统计
		if tCount[s[right]] == 0 {
			continue
		}

		windowCount[s[right]]++
		// 有可能有重复字符，不能重复减少 diff
		if windowCount[s[right]] <= tCount[s[right]] {
			diff--
		}

		for diff == 0 {
			if tCount[s[left]] == 0 {
				left++
				continue
			}

			curLen := right - left + 1
			if curLen < minLen {
				minLen = curLen
				minIdx = left
			}

			windowCount[s[left]]--
			if windowCount[s[left]] < tCount[s[left]] {
				diff++
			}

			left++
		}
	}

	if minIdx == -1 {
		return ""
	}
	return s[minIdx : minIdx+minLen]
}
