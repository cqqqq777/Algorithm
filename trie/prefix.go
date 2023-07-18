package trie

// FindPrefix 找到 arr2 以前缀身份出现在 arr1 中的字符串
func FindPrefix(arr1, arr2 []string) []string {
	root := NewTrie(arr1)
	result := make([]string, 0, len(arr2))
	for _, str := range arr2 {
		if PrefixNum(root, str) > 0 {
			result = append(result, str)
		}
	}
	return result
}

// PrefixNum 找到前缀树中以给定字符串作为前缀的数量
func PrefixNum(root *Node, prefix string) int {
	curNode := root
	for _, char := range prefix {
		index := int(char - 'a')
		if curNode.Nexts[index] == nil {
			return 0
		}
		curNode = curNode.Nexts[index]
	}
	return curNode.Path
}
