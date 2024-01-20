package day06

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	name  string
	nexts map[string]*TrieNode
}

// CatalogueStructure
// 给你一个字符串类型的数组arr，譬如：
// String[] arr = { "b/cst", "d/", "a/d/e", "a/b/c" };
// 你把这些路径中蕴含的目录结构给画出来，子目录直接列在父目录下面，并比父目录向右进两格，就像这样:
// 同一级的需要按字母顺序排列，不能乱。
func CatalogueStructure(catalogue []string) {
	// 构建前缀树
	root := &TrieNode{nexts: make(map[string]*TrieNode)}
	for _, path := range catalogue {
		cur := root
		for _, v := range strings.Split(path, "/") {
			if _, ok := cur.nexts[v]; !ok {
				cur.nexts[v] = &TrieNode{name: v, nexts: make(map[string]*TrieNode)}
			}
			cur = cur.nexts[v]
		}
	}
	// 深度优先遍历
	printDepth(root, 0)
}

func printDepth(cur *TrieNode, level int) {
	if level != 0 {
		space := getSpace(level)
		fmt.Println(space + cur.name)
	}
	for _, v := range cur.nexts {
		printDepth(v, level+1)
	}
}

func getSpace(level int) string {
	var b strings.Builder
	for i := 1; i < level; i++ {
		b.WriteString("  ")
	}
	return b.String()
}
