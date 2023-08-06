package day02

import (
	"fmt"
	"testing"
)

func TestMaxLength(t *testing.T) {
	for i := 0; i <= 50; i++ {
		if Winner(i) != table(i) {
			fmt.Println(i, false)
		}
	}
}
