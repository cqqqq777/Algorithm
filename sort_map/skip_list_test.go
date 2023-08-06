package sort_map

import (
	"fmt"
	"testing"
)

func TestSkipList(t *testing.T) {
	s := NewSkipList()
	for i := 1; i <= 100; i++ {
		s.Put(Key(i), Value(i))
	}
	fmt.Println(s.Get(34))
	fmt.Println(s.Get(77))
	fmt.Println(s.Get(55))
	s.Delete(55)
	fmt.Println(s.Get(55))
	fmt.Println(s.Get(56))
}
