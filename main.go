package main

import "fmt"

func main() {
	//nums := []int{1, 2, 3, 4, 5, 6, 4, 3, 2, 5, 6, 1, 1, 1, 2, 1, 1, 2, 3, 4, 4, 5}
	//
	//var eof int
	//for i := 0; i < len(nums); i++ {
	//	eof ^= nums[i] // eof = a ^ b
	//}
	//
	//rightOne := eof & (^eof + 1) // get the first num that is equal to 1
	//
	//var a, b int
	//for i := 0; i < len(nums); i++ {
	//	if rightOne&nums[i] == 0 {
	//		a ^= nums[i]
	//	}
	//}
	//
	//b = eof ^ a
	//
	//fmt.Println(a, b)

	//nums := []int{-1, 0, 3, 5, 9, 12, 14}
	//position := binarySearch(nums, 14)
	//fmt.Println(position)

	//nums := []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 6}
	//poi := search(nums, 2)
	//fmt.Println(poi)

	nums := []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6, 8, 8, 8}
	r := searchRange(nums, 7)
	fmt.Println(r)
}

func binarySearch(nums []int, target int) int {
	if target < nums[0] || target > nums[len(nums)-1] {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

func search(nums []int, target int) int {
	if target > nums[len(nums)-1] {
		return -1
	}

	left, right := 0, len(nums)-1
	position := 0

	for left <= right {
		mid := left + (right-left)>>2
		if nums[mid] >= target {
			position = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return position
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || target > nums[len(nums)-1] || target < nums[0] {
		return []int{-1, -1}
	}

	left, right := 0, len(nums)-1
	leftPosition, rightPosition := -1, -1

	for left <= right {
		mid := left + (right-left)>>2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			leftPosition = mid
			right = mid - 1
		}
	}

	left, right = 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			rightPosition = mid
			left = mid + 1
		}
	}

	return []int{leftPosition, rightPosition}
}
