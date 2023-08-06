package day01

// AppleMinBag 给定苹果数，返回最少袋子数
// 只有6和8两种袋子，如果凑不齐就返回-1
func AppleMinBag(num int) int {
	return appleMinBag1(num)
}

// 思路：优先使用8的袋子
func appleMinBag1(num int) int {
	if num < 6 || num == 7 || num&1 == 1 {
		return -1
	}
	min8, minBag := num/8, num/8
	rest := num % 8
	for rest != 0 {
		if rest%6 == 0 {
			minBag = min8 + rest/6
			break
		}
		min8--
		rest += 8
		// 当 rest 大于24时比不可能凑出来
		if rest > 24 {
			return -1
		}
	}
	return minBag
}

// 根据上面的贪心策略得出的结果，发现规律，直接按规律写代码
func appleMinBag2(num int) int {
	if num <= 8 {
		if num == 6 || num == 8 {
			return 1
		}
		return -1
	}
	if num&1 == 1 {
		return -1
	}
	return (num-1)/8 + 1
}
