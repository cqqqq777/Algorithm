package day03

import (
	"fmt"
	"testing"
)

func TestTopK(t *testing.T) {
	fmt.Println(TopK([]string{"az", "sf", "ff", "az", "az", "ii", "sf", "qd", "qd", "qd", "qd", "sd", "sf", "oi", "pa"}, 3))
}
