package day07

// ChineseExpress
// 把一个数字用中文表示出来。数字范围为 [0, 99999]。
// 使用数字取代中文数字注：对于 11 应该表示为 一十一，而不是十一
// 输入描述：
// 数字 0（包含）到 99999（包含）。
// 示例1:
// 输入
// 12001
// 输出
// 一万两千零一
func ChineseExpress(num int) string {
	if num == 0 {
		return "零"
	}
	res := ""
	if num < 0 {
		res += "负"
	}
	y := abs(num / 100000000)
	rest := abs(num % 100000000)
	if y == 0 {
		return res + num1To99999999(rest)
	}
	res += res + num1To9999(y) + "亿"
	if rest == 0 {
		return res
	}
	if rest < 10000000 {
		return res + "零" + num1To99999999(rest)
	}
	return res + num1To99999999(rest)
}

func num1To9(num int) string {
	if num < 1 || num > 9 {
		return ""
	}
	names := []string{"一", "二", "三", "四", "五", "六", "七", "八", "九"}
	return names[num-1]
}

func num1To99(num int) string {
	if num < 1 || num > 99 {
		return ""
	}
	if num < 10 {
		return num1To9(num)
	}
	highBit := num / 10
	return num1To9(highBit) + "十" + num1To9(num%10)
}

func num1To999(num int) string {
	if num < 1 || num > 999 {
		return ""
	}
	if num < 100 {
		return num1To99(num)
	}
	res := num1To9(num/100) + "百"
	rest := num % 100
	if rest == 0 {
		return res
	} else if rest < 10 {
		return res + "零" + num1To9(rest)
	}
	return res + num1To99(rest)
}

func num1To9999(num int) string {
	if num < 1 || num > 9999 {
		return ""
	}
	if num < 1000 {
		return num1To999(num)
	}
	res := num1To9(num/1000) + "千"
	rest := num % 1000
	if rest == 0 {
		return res
	} else if rest < 100 {
		return res + "零" + num1To99(rest)
	}
	return res + num1To999(rest)
}

func num1To99999999(num int) string {
	if num < 1 || num > 99999999 {
		return ""
	}
	if num < 10000 {
		return num1To9999(num)
	}
	res := num1To9999(num/10000) + "万"
	rest := num % 10000
	if rest == 0 {
		return res
	} else if rest < 1000 {
		return res + "零" + num1To999(rest)
	}
	return res + num1To9999(rest)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
