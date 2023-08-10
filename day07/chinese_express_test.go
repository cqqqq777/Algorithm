package day07

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestChineseExpress(t *testing.T) {
	for i := 0; i < 100; i++ {
		r := rand.Intn(10000000000)
		fmt.Printf("%d %s\n", r, ChineseExpress(r))
	}
}
