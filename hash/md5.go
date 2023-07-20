package hash

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5() {
	m := md5.New()
	io.WriteString(m, "hello")
	fmt.Printf("%x\n", m.Sum(nil))
	m.Reset()
	fmt.Printf("%x", m.Sum(nil))
}
