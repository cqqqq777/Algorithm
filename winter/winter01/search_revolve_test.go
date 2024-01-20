package winter01

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	res := search([]int{3, 1}, 1)
	fmt.Println(res)
}
