package solution

import (
	"math"
	"strings"
)

//先处理掉一些非法数值，留下纯数字
//接着对纯数字字符串进行处理，判断是否溢出
func myAtoi(s string) int {
	return convert(clean(s))
}

//处理字符串
func clean(s string) (sign int, nums string) {
	//去除首尾空格
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}
	switch s[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, nums = 1, s
	case '+':
		sign, nums = 1, s[1:]
	case '-':
		sign, nums = -1, s[1:]
	default:
		nums = ""
		return
	}
	//取前i个数字字符
	for i, num := range nums {
		if num < '0' || num > '9' {
			nums = nums[:i]
			break
		}
	}
	return
}

//返回数字
func convert(sign int, nums string) int {
	ans := 0
	for _, num := range nums {
		//把字符类型转化为int
		ans = ans*10 + int(num-'0')
		//检查溢出
		switch {
		case sign == 1 && ans > math.MaxInt32:
			return math.MaxInt32
		case sign == -1 && -ans < math.MinInt32:
			return math.MinInt32
		}
	}
	return sign * ans
}
