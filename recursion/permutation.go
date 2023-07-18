package recursion

import "fmt"

// Permutation 打印一个字符串的全排列，不出现重复
func Permutation(str string) {
	if len(str) == 0 {
		return
	}
	bytes := make([]byte, len(str))
	copy(bytes, str)
	PermutationRecur(bytes, 0)
}

// PermutationRecur 递归
// index 是即将要被选择的字符
// [i..] 范围上的所有字符都可以被选择
// [0..i-1] 是前面已经选择出来的结果
func PermutationRecur(bytes []byte, index int) {
	if index == len(bytes) {
		fmt.Println(string(bytes))
		return
	}
	// 使用 map 存储当前位置被选择的字符，如果有重复的就不选，达到去重的目的
	visited := make(map[byte]struct{})
	for j := index; j < len(bytes); j++ {
		if _, ok := visited[bytes[j]]; !ok {
			visited[bytes[j]] = struct{}{}
			// 交换位置即代表 index 位置被选择，然后进行选择下一个位置的字符
			bytes[index], bytes[j] = bytes[j], bytes[index]
			PermutationRecur(bytes, index+1)
			// 还原后以进行下一次对 index 位置的选择
			bytes[index], bytes[j] = bytes[j], bytes[index]
		}
	}
}
