package trie

// CommonString 找到在 arr2 中同时也出现在 arr1 中的字符串
func CommonString(arr1, arr2 []string) []string {
	root := NewTrie(arr1)
	result := make([]string, 0, len(arr2))
	for _, str := range arr2 {
		if _, exist := SearchString(root, str); exist {
			result = append(result, str)
		}
	}
	return result
}
