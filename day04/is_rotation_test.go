package day04

import (
	"fmt"
	"testing"
)

func TestIsRotation(t *testing.T) {
	fmt.Println(IsRotation("abcd", "cdab"))                                 //expected output: true
	fmt.Println(IsRotation("akjdhgjkasgfdjkahsdk", "kahsdkakjdhgjkasgfdj")) // expected output: true
	fmt.Println(IsRotation("aklsdhgask", "alskhdkogiyhig"))                 // expected output: false
}
