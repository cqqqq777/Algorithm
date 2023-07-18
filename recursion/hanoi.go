package recursion

import "fmt"

// Hanoi 汉诺塔问题
func Hanoi(n int) {
	if n <= 0 {
		return
	}
	HanoiRecur(n, n, "left", "right", "middle")
}

func HanoiRecur(rest, down int, from, to, help string) {
	if rest == 1 {
		fmt.Println("move", down, "from", from, "to", to)
	} else {
		HanoiRecur(rest-1, down-1, from, help, to)
		HanoiRecur(1, down, from, to, help)
		HanoiRecur(rest-1, down-1, help, to, from)
	}

}
