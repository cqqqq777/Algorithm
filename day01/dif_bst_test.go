package day01

import (
	"fmt"
	"testing"
	"time"
)

func TestDifBST(t *testing.T) {
	start := time.Now()
	fmt.Println(DifBST(1500))
	fmt.Println(time.Since(start).Milliseconds())
	//start = time.Now()
	//fmt.Println(process1(20))
	//fmt.Println(time.Since(start).Milliseconds())
	start = time.Now()
	fmt.Println(process3(1500))
	fmt.Println(time.Since(start).Milliseconds())
}
