package median_finder

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	mf := Constructor()
	mf.AddNum(6)
	fmt.Println(mf.FindMedian())
	mf.AddNum(10)
	fmt.Println(mf.FindMedian())
	mf.AddNum(2)
	fmt.Println(mf.FindMedian())
	mf.AddNum(6)
	fmt.Println(mf.FindMedian())

}
