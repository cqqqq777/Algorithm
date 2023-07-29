package bitwise_operate

import (
	"fmt"
	"testing"
)

func TestGetMax(t *testing.T) {
	fmt.Println(GetMax(4, 19874))
}

func TestIs2Power(t *testing.T) {
	fmt.Println(Is2Power(1024)) // expected output: true
	fmt.Println(Is2Power(1023)) // expected output: false
}

func TestIs4Power(t *testing.T) {
	fmt.Println(Is4Power(16))  // expected output: true
	fmt.Println(Is4Power(64))  // expected output: true
	fmt.Println(Is4Power(328)) // expected output: false
	//fmt.Printf("%b", 0x55555555<<1)
}

func TestPlus(t *testing.T) {
	fmt.Println(Plus(100, 200))     // expected output: 300
	fmt.Println(Subtract(100, 100)) // expected output: 0
	fmt.Println(Multiply(23, 2))    // expected output: 46
	fmt.Println(Divide(14, 9))      // expected output: 4
}
