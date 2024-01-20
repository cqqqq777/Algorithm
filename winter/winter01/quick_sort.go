package winter01

import "math/rand"

// QuickSort 3.0 版本，每次随机选取一个数作为基准
func QuickSort(sli []int) {
	quickSort(sli, 0, len(sli)-1)
}

func quickSort(arr []int, l, r int) {
	if l < r {
		randIndex := rand.Intn(r-l+1) + l
		arr[randIndex], arr[r] = arr[r], arr[randIndex]
		p := partition(arr, l, r)
		quickSort(arr, l, p[0]-1)
		quickSort(arr, p[1]+1, r)
	}
}

// 返回小于区的右边界和大于区的左边界
func partition(arr []int, l, r int) []int {
	less, more := l-1, r
	for l < more {
		if arr[l] < arr[r] {
			arr[l], arr[less+1] = arr[less+1], arr[l]
			less++
			l++
		} else if arr[l] > arr[r] {
			arr[l], arr[more-1] = arr[more-1], arr[l]
			more--
		} else {
			l++
		}
	}
	arr[r], arr[more] = arr[more], arr[r]
	return []int{less + 1, more}
}
