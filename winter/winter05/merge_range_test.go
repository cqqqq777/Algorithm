package winter05

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	res := longestCommonSubsequence("abc", "akboc")
	fmt.Println(res)
}
