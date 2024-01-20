package str

// KMP returns the index of the first instance of son in father, or -1 if father is not present in son.
func KMP(father, son string) int {
	if len(father) < len(son) || len(father) == 0 || len(son) == 0 {
		return -1
	}
	fatherStart, sonStart := 0, 0
	next := GetNext(son)
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

func GetNext(str string) []int {
	if len(str) <= 0 {
		return []int{}
	}
	result := make([]int, len(str))
	result[0] = -1

	j := 0
	for i := 2; i < len(str); {
		if str[j] == str[i-1] {
			result[i] = j + 1
			i++
			j++
		} else if j == 0 {
			result[i] = 0
			i++
		} else {
			j = result[j]
		}
	}
	return result
}
