package day06

import (
	"fmt"
	"testing"
)

func TestRecentMaxScore(t *testing.T) {
	fmt.Println(RecentMaxScore([]int{1, 1, -1, -10, 11, 4, -6, 9, 20, -10, -2})) // expected output: 38
}
