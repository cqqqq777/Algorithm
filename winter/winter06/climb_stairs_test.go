package winter06

import (
	"fmt"
	"testing"
)

func TestClimb(t *testing.T) {
	for i := 1; i <= 45; i++ {
		fmt.Println(climbStairs(i))
	}
}
