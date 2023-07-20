package hash

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	r := NewRandomPool()
	r.Insert("A")
	r.Insert("b")
	r.Insert("a")
	r.Insert("ff")
	r.Insert("aq")
	fmt.Println(r.GetRandom())
	r.Delete("a")
	fmt.Println(r.GetRandom())
	fmt.Println(r.GetRandom())
	fmt.Println(r.GetRandom())
	fmt.Println(r.GetRandom())
	fmt.Println(r.GetRandom())
}
