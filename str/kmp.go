package str

// KMP returns the index of the first instance of son in father, or -1 if father is not present in son.
func KMP(father, son string) int {
	if len(father) < len(son) || len(father) == 0 || len(son) == 0 {
		return -1
	}
	fatherStart, sonStart := 0, 0
	next := GetNexts(son)
	for fatherStart < len(father) && sonStart < len(son) {
		if father[fatherStart] == son[sonStart] {
			fatherStart++
			sonStart++
		} else if sonStart == 0 {
			fatherStart++
		} else {
			sonStart = next[sonStart]
		}
	}
	if sonStart == len(son) {
		return fatherStart - sonStart
	}
	return -1
}

func GetNexts(str string) []int {
	if len(str) <= 0 {
		return []int{}
	}
	result := make([]int, len(str))
	result[0] = -1
	for i := 2; i < len(str); i++ {
		for j := 1; j < i; j++ {
			if str[:j] == str[i-j:i] {
				result[i] = j
			}
		}
	}
	return result
}
