package recursion

// Convert 返回有多少种将 arr 转化为字符串的转化结果
func Convert(arr []int) int {
	if arr == nil {
		return 0
	}
	return ConvertRecur(arr, 0)
}

func ConvertRecur(arr []int, index int) int {
	if index >= len(arr)-1 {
		return 1
	}
	res := 0
	if arr[index] == 1 || (arr[index] == 2 && arr[index+1] <= 6) {
		res += ConvertRecur(arr, index+2)
	}
	res += ConvertRecur(arr, index+1)
	return res
}
