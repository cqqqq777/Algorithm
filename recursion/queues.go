package recursion

import "math"

// NQueues n皇后问题
func NQueues(n int) int {
	if n <= 3 {
		return 0
	}
	// record used to store the number of columns that have been placed
	record := make([]int, n)
	return Num1(record, 0, n)
}

func Num1(record []int, i, n int) int {
	if i == n {
		return 1
	}
	res := 0
	for j := 0; j < n; j++ {
		if isValid(record, i, j) {
			record[i] = j
			res += Num1(record, i+1, n)
		}
	}
	return res
}

func isValid(record []int, i, j int) bool {
	for k := 0; k < i; k++ {
		if record[k] == j || math.Abs(float64(i-k)) == math.Abs(float64(j-record[k])) {
			return false
		}
	}
	return true
}

// NQueues2 n皇后问题，只能处理64位以内的问题
func NQueues2(n int) int {
	if n <= 3 || n >= 64 {
		return 0
	}
	var limit int64 = (1 << n) - 1
	return Num2(limit, 0, 0, 0)
}

func Num2(limit int64, leftLim, rightLim, colLim int64) int {
	if colLim == limit {
		return 1
	}
	res := 0
	pos := limit & (^(leftLim | rightLim | colLim))
	for pos != 0 {
		mostRight := pos & (^pos + 1)
		pos = pos - mostRight
		res += Num2(limit,
			(leftLim+mostRight)<<1,
			(rightLim+mostRight)>>1,
			colLim+mostRight)
	}
	return res
}
