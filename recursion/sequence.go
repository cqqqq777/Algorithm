package recursion

import "fmt"

// Sequence 打印一个字符串的所有子序列
func Sequence(str string) {
	if len(str) == 0 {
		return
	}
	bytes := make([]byte, len(str))
	copy(bytes, str)
	SequenceRecur(bytes, 0)
}

func SequenceRecur(bytes []byte, index int) {
	if index == len(bytes) {
		fmt.Println(string(bytes))
		return
	}
	SequenceRecur(bytes, index+1)
	tmp := bytes[index]
	bytes[index] = 0
	SequenceRecur(bytes, index+1)
	bytes[index] = tmp
}
