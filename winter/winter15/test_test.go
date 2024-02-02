package winter15

import (
	"runtime"
	"testing"
)

func Test(t *testing.T) {
	s := make([]int, 10, 11)
	t.Logf("s: %p, len: %d, cap: %d", s, len(s), cap(s))
	s1 := append(s, 1)
	t.Logf("s: %p, len: %d, cap: %d", s1, len(s1), cap(s1))
	t.Log(s1)
	s = append(s, 2)
	t.Log(s1)
	runtime.Gosched()
}
