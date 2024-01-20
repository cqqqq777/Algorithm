package recursion

// MaxValue 解决 knapsack 问题
func MaxValue(weights, values []int, bag int) int {
	if weights == nil || values == nil || len(weights) != len(values) || bag < 1 {
		return 0
	}
	return MaxValueRecur(weights, values, bag, 0, 0, 0)
}

func MaxValueRecur(weights, values []int, bag, weight, value, index int) int {
	if index == len(weights)-1 {
		if weight+weights[index] <= bag {
			return value + values[index]
		}
		return value
	}
	result := 0
	if weights[index]+weight <= bag {
		result = MaxValueRecur(weights, values, bag, weight+weights[index], value+values[index], index+1)
	}
	notChoose := MaxValueRecur(weights, values, bag, weight, value, index+1)
	if result < notChoose {
		result = notChoose
	}
	return result
}
