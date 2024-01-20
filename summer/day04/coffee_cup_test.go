package day04

import (
	"fmt"
	"testing"
	"time"
)

func TestWashCupMinTime(t *testing.T) {
	start := time.Now()
	fmt.Println(WashCupMinTime([]int{1, 2, 3, 4, 5}, 20, 1, 5))
	fmt.Println(time.Since(start).Milliseconds())
	start = time.Now()
	fmt.Println(WashCupMinTime2([]int{1, 2, 3, 4, 5}, 20, 1, 5))
	fmt.Println(time.Since(start).Milliseconds())

}
