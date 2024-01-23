package winter05

import (
	"strconv"
	"strings"
)

// 复原 ip 地址
// 返回所有有效 ip
func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return []string{}
	}

	track := make([]string, 0, 4)
	res := make([]string, 0)

	var recur func(str []string, idx int)
	recur = func(str []string, idx int) {
		if len(str) == 4 && idx == len(s) {
			res = append(res, strings.Join(str, "."))
			return
		}
		if len(str) == 4 && idx < len(s) {
			return
		}
		for i := 1; i <= 3; i++ {
			if idx+i > len(s) {
				return
			}
			if s[idx] == '0' && i != 1 {
				return
			}
			choice := s[idx : idx+i]
			if n, _ := strconv.Atoi(choice); n > 255 {
				return
			}
			str = append(str, choice)
			recur(str, idx+i)
			str = str[:len(str)-1]
		}
	}
	recur(track, 0)
	return res
}
