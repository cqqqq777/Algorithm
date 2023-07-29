package recursion

import (
	"fmt"
	"testing"
	"time"
)

func TestWinnerScore(t *testing.T) {
	start := time.Now()
	fmt.Println(WinnerScore([]int{12, 2, 3, 4, 5, 2, 1, 3, 5, 76, 2, 4, 356464, 233, 33, 12, 3, 5, 7, 4, 12, 2, 344, 6, 7, 2, 21, 24, 6, 4, 234, 444, 4}))
	fmt.Println(time.Since(start).Microseconds())
	start = time.Now()
	fmt.Println(WinnerScoreDP([]int{12, 2, 3, 4, 5, 2, 1, 3, 5, 76, 2, 4, 356464, 233, 33, 12, 3, 5, 7, 4, 12, 2, 344, 6, 7, 2, 21, 24, 6, 4, 234, 444, 4}))
	fmt.Println(time.Since(start).Microseconds())
}
