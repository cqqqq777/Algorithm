package day01

import "math/rand"

type RandFunc func() int

// Rand5To7 给定随机函数能生成1-5，要求使其能随机生成1-7
// 利用给定函数，随机生成的数里面大于3的就当作1，小于三的当作0，等于3就重新生成
// 生成三次分别作为二进制的最后三位，如果得到结果为7就重新生成,得到结果加一返回即可
func Rand5To7(f RandFunc) int {
	res := 7
	for res == 7 {
		res = rand01(1, 5, f)<<0 + rand01(1, 5, f)<<1 + rand01(1, 5, f)<<2
	}
	return res + 1
}

func rand01(start, stop int, f RandFunc) int {
	if (stop-start)&1 == 0 {
		mid := start + (stop-start)>>1
		r := f()
		if r > mid {
			return 1
		}
		return 0
	} else {
		mid := start + (stop-start)>>1
	lab:
		r := f()
		if r > mid {
			return 1
		} else if r < mid {
			return 0
		}
		goto lab
	}
}

func rand5() int {
	return rand.Intn(5) + 1
}

// Rand01 等概率返回0或1
// 给定函数 f，p 概率返回0，1-p 概率返回1，
func Rand01(f RandFunc) int {
	res := 0
	for res != 1 && res != 2 {
		res = f()<<1 + f()<<0
	}
	return res - 1
}

// 70% 返回0，30% 返回1
func r() int {
	if rand.Intn(100) >= 70 {
		return 0
	}
	return 1
}
