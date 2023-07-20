package union_find_set

import (
	"fmt"
	"testing"
)

func TestUnionFindSet(t *testing.T) {
	elements := make([]*Element, 100)
	for i := 0; i < 100; i++ {
		elements[i] = &Element{Value: i}
	}
	usf := NewUnionFindSet(elements...)
	fmt.Println(usf.IsSameSet(elements[0], elements[1])) // expected output: false
	usf.Union(elements[0], elements[1])
	fmt.Println(usf.IsSameSet(elements[0], elements[1])) // expected output: true
}
