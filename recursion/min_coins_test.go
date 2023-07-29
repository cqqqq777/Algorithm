package recursion

import (
	"fmt"
	"testing"
	"time"
)

func TestMinCoins(t *testing.T) {
	start := time.Now()
	fmt.Println(MinCoins([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 10000}, 10000))
	fmt.Println(time.Since(start).Microseconds())
	start = time.Now()
	fmt.Println(MinCoins2([]int{10}, 1))
	fmt.Println(time.Since(start).Microseconds())
	start = time.Now()
	fmt.Println(MinCoins3([]int{1}, 1))
	fmt.Println(time.Since(start).Microseconds())
}
