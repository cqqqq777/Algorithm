package recursion

import (
	"fmt"
	"testing"
	"time"
)

func TestNQueues(t *testing.T) {
	start := time.Now()
	fmt.Println(NQueues(14))                      // output: 365596
	fmt.Println(time.Since(start).Milliseconds()) // output: 4432
	start = time.Now()
	fmt.Println(NQueues2(14))                     // output: 365596
	fmt.Println(time.Since(start).Milliseconds()) // output: 144
}
