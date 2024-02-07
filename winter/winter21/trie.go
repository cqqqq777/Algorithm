package winter21

// Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
// 请你实现 Trie 类：
// Trie() 初始化前缀树对象。
// void insert(String word) 向前缀树中插入字符串 word 。
// boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
// boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

type Trie struct {
	nexts     []*Trie
	pass, end int
}

func Constructor() Trie {
	return Trie{
		nexts: make([]*Trie, 26),
		pass:  0,
		end:   0,
	}
}

func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		tmp := cur.nexts[word[i]-'a']
		if tmp == nil {
			tmp = &Trie{
				nexts: make([]*Trie, 26),
				pass:  0,
				end:   0,
			}
			cur.nexts[word[i]-'a'] = tmp
		}
		tmp.pass++
		if i == len(word)-1 {
			tmp.end++
		}
		cur = tmp
	}
}

func (this *Trie) Search(word string) bool {
	cur := this
	for i := 0; i < len(word); i++ {
		tmp := cur.nexts[word[i]-'a']
		if tmp == nil {
			return false
		}
		if i == len(word)-1 && tmp.end > 0 {
			return true
		}
		cur = tmp
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix); i++ {
		tmp := cur.nexts[prefix[i]-'a']
		if tmp == nil {
			return false
		}
		cur = tmp
	}
	return true
}
