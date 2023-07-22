package binary_tree

import (
	"fmt"
	"testing"
)

func TestMaxHappy(t *testing.T) {
	boss := &Employee{happy: 0, subordinates: make([]*Employee, 3)}
	for i := 0; i < 3; i++ {
		boss.subordinates[i] = &Employee{happy: 190, subordinates: make([]*Employee, 1)}
		boss.subordinates[i].subordinates[0] = &Employee{happy: 9, subordinates: nil}
	}
	fmt.Println(MaxHappy(boss))
}
