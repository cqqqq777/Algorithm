package dp

import (
	"fmt"
	"testing"
	"time"
)

func TestHorseJump(t *testing.T) {
	start := time.Now()
	fmt.Println(HorseJump(6, 6, 12))
	fmt.Println(time.Since(start).Milliseconds())
	start = time.Now()
	fmt.Println(horseJumpDP(6, 6, 12))
	fmt.Println(time.Since(start).Milliseconds())
}
