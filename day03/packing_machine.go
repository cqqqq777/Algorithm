package day03

func PackingMachine(arr []int) int {
	sum, size := 0, len(arr)
	for _, v := range arr {
		sum += v
	}
	if sum%size != 0 {
		return -1
	}
	res, avg := 0, sum/size
	leftSum := 0
	for i := 0; i < size; i++ {
		l := leftSum - i*avg
		r := (sum - leftSum - arr[i]) - (size-i-1)*avg
		option := ops(l, r)
		if option > res {
			res = option
		}
		leftSum += arr[i]
	}
	return res
}

func ops(left, right int) int {
	if left < 0 && right < 0 {
		return -(left + right)
	}
	return Max(Abs(left), Abs(right))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
