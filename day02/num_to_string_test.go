package day02

import (
	"fmt"
	"testing"
	"time"
)

func TestNumToString(t *testing.T) {
	start := time.Now()
	fmt.Println(NumToString(1276223617))
	fmt.Println(time.Since(start).Milliseconds())
	start = time.Now()
	fmt.Println(NumToString2(1276223617))
	fmt.Println(time.Since(start).Milliseconds())
}
