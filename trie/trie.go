package trie

type Node struct {
	Path, End int
	Nexts     []*Node
}

// NewTrie 新建一颗前缀树，其中的每个字符串的字符都在 a-z 26个字符之间
func NewTrie(strings []string) *Node {
	root := &Node{Nexts: make([]*Node, 26)}
	for _, str := range strings {
		curNode := root
		curNode.Path++
		index := 0
		for _, char := range str {
			index = int(char - 'a')
			if curNode.Nexts[index] == nil {
				curNode.Nexts[index] = &Node{Nexts: make([]*Node, 26)}
			}
			curNode = curNode.Nexts[index]
			curNode.Path++
		}
		curNode.End++
	}
	return root
}

// Insert 向前缀树插入字符串，如果 root 为空则创建一个前缀树
func Insert(root *Node, strings ...string) *Node {
	if root == nil {
		root = NewTrie(strings)
		return root
	}
	for _, str := range strings {
		curNode := root
		curNode.Path++
		index := 0
		for _, char := range str {
			index = int(char - 'a')
			if curNode.Nexts[index] == nil {
				curNode.Nexts[index] = &Node{Nexts: make([]*Node, 26)}
			}
			curNode = curNode.Nexts[index]
			curNode.Path++
		}
		curNode.End++
	}
	return root
}

// SearchString 在前缀树中寻找 string，如果存在同时返回其个数
func SearchString(root *Node, str string) (int, bool) {
	if root == nil {
		return 0, false
	}
	curNode := root
	for _, char := range str {
		index := int(char - 'a')
		if curNode.Nexts[index] == nil {
			return 0, false
		}
		curNode = curNode.Nexts[index]
	}
	if curNode.End == 0 {
		return 0, false
	}
	return curNode.End, true
}

// Delete 删除前缀树中的某一字符串
func Delete(root *Node, str string) {
	_, ok := SearchString(root, str)
	if !ok {
		return
	}
	curNode := root
	curNode.Path--
	for _, char := range str {
		index := int(char - 'a')
		if curNode.Nexts[index].Path-1 == 0 {
			curNode.Nexts[index] = nil
			return
		}
		curNode.Nexts[index].Path--
		curNode = curNode.Nexts[index]
	}
	curNode.End--
}
