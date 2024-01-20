package day01

import (
	"fmt"
	"testing"
)

func TestAppleMinBag(t *testing.T) {
	for i := 0; i <= 100; i++ {
		if appleMinBag1(i) != appleMinBag2(i) {
			fmt.Println(false)
			return
		}
	}
	fmt.Println(true)
}
