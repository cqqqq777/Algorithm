package winter14

import (
	"strconv"
	"strings"
)

// 验证是否是合法的 ipv4 或者 ipv6 地址
func validIPAddress(queryIP string) string {
	if ip := strings.Split(queryIP, "."); len(ip) == 4 {
		for _, v := range ip {
			if len(v) > 1 && v[0] == '0' {
				return "Neither"
			}
			if num, err := strconv.Atoi(v); err != nil || num > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	}
	if ip := strings.Split(queryIP, ":"); len(ip) == 8 {
		for _, v := range ip {
			if len(v) > 1 && v[0] == '0' {
				return "Neither"
			}
			if len(v) > 4 {
				return "Neither"
			}
			if _, err := strconv.ParseUint(v, 16, 64); err != nil {
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}
