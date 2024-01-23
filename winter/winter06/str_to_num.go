package winter06

// TODO
func myAtoi(s string) int {
	if s == "" {
		return 0
	}

	flag := true

	for _, v := range s {
		if flag && v == ' ' {
			continue
		}
		if flag && v == '-' {
			continue
		}
		flag = false

	}
	return 9
}
