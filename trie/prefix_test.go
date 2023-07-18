package trie

import (
	"fmt"
	"testing"
)

func TestPrefixNum(t *testing.T) {
	root := NewTrie([]string{"abc", "a", "bc", "abd", "abc"})
	fmt.Println(PrefixNum(root, "ab")) // expected output: 3
}

func TestFindPrefix(t *testing.T) {
	arr1 := []string{"abc", "a", "bc", "abd", "abc"}
	arr2 := []string{"ab", "a", "bc", "ad"}
	fmt.Println(FindPrefix(arr1, arr2)) // expected output: ["ab", "a", "bc"]
}
