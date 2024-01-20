package day07

import (
	"fmt"
	"testing"
	"time"
)

func TestKiki(t *testing.T) {
	s := time.Now()
	fmt.Println(Kiki(6, 44, 2, 2, 2))
	fmt.Println(time.Since(s).Microseconds())
	s = time.Now()
	fmt.Println(minCoinDP(6, 44, 2, 2, 2))
	fmt.Println(time.Since(s).Microseconds())
}
