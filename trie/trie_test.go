package trie

import (
	"fmt"
	"testing"
)

func TestNewTrie(t *testing.T) {
	root := NewTrie([]string{"abc", "a", "bc", "abd"})
	fmt.Println(root.Path)                           // expected output: 4
	fmt.Println(root.Nexts[0].Path)                  // expected output: 3
	fmt.Println(root.Nexts[0].Nexts[1].Path)         // expected output: 2
	fmt.Println(root.Nexts[0].End)                   // expected output: 1
	fmt.Println(root.Nexts[1].Path)                  // expected output: 1
	fmt.Println(root.Nexts[2])                       // expected output: nil
	fmt.Println(root.Nexts[0].Nexts[1].Nexts[3].End) // expected output: 1
}

func TestSearchString(t *testing.T) {
	root := NewTrie([]string{"abc", "a", "bc", "abd", "abc"})
	fmt.Println(SearchString(root, "abc"))  // expected output: 2 true
	fmt.Println(SearchString(root, "abcd")) // expected output: 0,false
	fmt.Println(SearchString(root, "ab"))   // expected output: 0,false
}

func TestInsert(t *testing.T) {
	root := NewTrie([]string{"abc", "a", "bc", "abd"})
	Insert(root, "abc", "a")
	fmt.Println(root.Nexts[0].End)                   // expected output:2
	fmt.Println(root.Nexts[0].Nexts[1].Nexts[2].End) // expected output: 2
}

func TestDelete(t *testing.T) {
	strings := []string{"abc", "a", "bc", "abd"}
	root := NewTrie(strings)
	fmt.Println(root)
	Delete(root, "abc")
	Delete(root, "a")
	Delete(root, "bc")
	Delete(root, "abd")
	fmt.Println(root)
}
