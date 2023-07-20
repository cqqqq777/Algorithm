package str

import (
	"fmt"
	"testing"
)

func TestKMP(t *testing.T) {
	fmt.Println(KMP("bbbbba", "b"))                            // expected output: 0
	fmt.Println(KMP("jashgdkjgfhaskjdg", "gdk"))               // expected output: 4
	fmt.Println(KMP("ADHSGTKJFASDHGKJAGHKJRTQGHKJWHGFK", "s")) // expected output: -1
}
