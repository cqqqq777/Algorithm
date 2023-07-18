package sort

import "fmt"

// MinSum 小和问题
func MinSum(sli []int) int {
	l, r := 0, len(sli)-1
	return Pro(sli, l, r)
}

func Pro(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	mid := l + (r-l)>>1
	return Pro(arr, l, mid) + Pro(arr, mid+1, r) + MinSumMerge(arr, l, r)
}

func MinSumMerge(arr []int, l, r int) (sum int) {
	help := make([]int, r-l+1)
	mid := l + (r-l)>>1
	i, p1, p2 := 0, l, mid+1
	for p1 <= mid && p2 <= r {
		if arr[p1] < arr[p2] {
			sum += arr[p1] * (r - p2 + 1)
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
		i++
	}

	for p1 <= mid {
		help[i] = arr[p1]
		i++
		p1++
	}
	for p2 <= r {
		help[i] = arr[p2]
		i++
		p2++
	}
	i = 0
	for l <= r {
		arr[l] = help[i]
		i++
		l++
	}
	return sum
}

// ReverseOrderPair  获取数组中所有逆序对
func ReverseOrderPair(sli []int) [][]int {
	l, r := 0, len(sli)-1
	return ReversePro(sli, l, r)
}

func ReversePro(sli []int, l, r int) [][]int {
	if l == r {
		return nil
	}
	mid := l + (r-l)>>1
	result := make([][]int, 0)
	result = append(result, ReversePro(sli, l, mid)...)
	result = append(result, ReversePro(sli, mid+1, r)...)
	result = append(result, ReverseMerge(sli, l, r)...)
	return result
}

func ReverseMerge(arr []int, l, r int) [][]int {
	help := make([]int, r-l+1)
	result := make([][]int, 0)
	mid := l + (r-l)>>1
	i, p1, p2 := 0, l, mid+1
	for p1 <= mid && p2 <= r {
		if arr[p1] < arr[p2] {
			help[i] = arr[p1]
			p1++
		} else {
			if arr[p1] >= arr[p2] {
				for j := p1; j <= mid; j++ {
					if arr[j] > arr[p2] {
						result = append(result, []int{arr[j], arr[p2]})
						fmt.Println(arr[j], arr[p2])
					}
				}
			}
			help[i] = arr[p2]
			p2++
		}
		i++
	}

	for p1 <= mid {
		help[i] = arr[p1]
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

// Partition1 给定一个 num 和 arr ， 将小于等于 num 的放在左边，大于 num 的放右边
func Partition1(sli []int, num int) {
	less := -1
	for i := 0; i < len(sli); i++ {
		if sli[i] <= num {
			sli[less+1], sli[i] = sli[i], sli[less+1]
			less++
		}
	}
}

// Partition2 给定一个 num 和 arr ，将小于 num 的放左边，大于 num 的放右边，等于 num 的放中间
func Partition2(sli []int, num int) {
	less, more := -1, len(sli)
	for i := 0; i < more; i++ {
		if sli[i] < num {
			sli[less+1], sli[i] = sli[i], sli[less+1]
			less++
		} else if sli[i] > num {
			sli[more-1], sli[i] = sli[i], sli[more-1]
			more--
			i--
		}
	}
}
