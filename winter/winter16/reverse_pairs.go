package winter16

// 在股票交易中，如果前一天的股价高于后一天的股价，则可以认为存在一个「交易逆序对」。
// 请设计一个程序，输入一段时间内的股票交易记录 record，返回其中存在的「交易逆序对」总数。
func reversePairs(record []int) int {
	// 利用归并排序
	if len(record) <= 1 {
		return 0
	}
	l, r := 0, len(record)-1
	return ReversePro(record, l, r)
}

func ReversePro(sli []int, l, r int) int {
	if l == r {
		return 0
	}
	mid := l + (r-l)>>1
	var result int
	result += ReversePro(sli, l, mid)
	result += ReversePro(sli, mid+1, r)
	result += ReverseMerge(sli, l, r)
	return result
}

func ReverseMerge(arr []int, l, r int) int {
	help := make([]int, r-l+1)
	var result int
	mid := l + (r-l)>>1
	i, p1, p2 := 0, l, mid+1
	for p1 <= mid && p2 <= r {
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			result += p2 - (mid + 1)
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
		i++
	}

	for p1 <= mid {
		help[i] = arr[p1]
		result += r - (mid + 1) + 1
		p1++
		i++
	}
	for p2 <= r {
		help[i] = arr[p2]
		p2++
		i++
	}
	for i = 0; l <= r; l++ {
		arr[l] = help[i]
		i++
	}
	return result
}
