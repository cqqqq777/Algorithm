package dp

import (
	"fmt"
	"testing"
)

func TestBobDie(t *testing.T) {
	fmt.Println(BobDie(10, 10, 3, 2, 5))
	fmt.Println(BobDieDP(10, 10, 3, 2, 5))
}
