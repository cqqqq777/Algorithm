package trie

import (
	"fmt"
	"testing"
)

func TestCommonString(t *testing.T) {
	result := CommonString([]string{"abc", "adv", "ssf", "aqwe"}, []string{"abc", "sss", "q", "adv"})
	for _, v := range result {
		fmt.Printf("%s ", v) // expected output: abc adv
	}
}
