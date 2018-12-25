package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/**
题目描述
	将一个字符串转换成一个整数(实现Integer.valueOf(string)的功能，
	但是string不符合数字要求时返回0)，要求不能使用字符串转换整数的库函数。
	数值为0或者字符串不是一个合法的数值则返回0。
输入描述:
	输入一个字符串,包括数字字母符号,可以为空
输出描述:
	如果是合法的数值表达则返回该数字，否则返回0

示例1
输入
+2147483647
1a33
输出
2147483647
0
*/
func main() {
	fmt.Println(StrToInt("12345"))
	fmt.Println(StrToInt("+12345"))
	fmt.Println(StrToInt("-12345"))
	fmt.Println(StrToInt("12345123451234512345123451234512345123451234512345"))
	fmt.Println(StrToInt("1234a5"))
	fmt.Println(StrToInt("12345.542"))

	fmt.Println(strconv.Atoi("12345.808"))
}

func StrToInt(str string) (result int) {
	str = strings.TrimSpace(str)
	if str == "" {
		return
	}
	// 正负号
	i := 1
	if strings.HasPrefix(str, "+") {
		i = 1
		str = strings.TrimLeft(str, "+")
	} else if strings.HasPrefix(str, "-") {
		i = -1
		str = strings.TrimLeft(str, "-")
	}

	//循环每个数字
	for _, v := range []rune(str) {
		if '0' <= v && '9' >= v {
			// 0-9为合法

			num := int(v - '0')

			// 溢出
			if result > math.MaxInt32-num {
				return 0
			}

			result = result*10 + num

			// 溢出
			if result < 0 {
				return 0
			}

		} else {
			// 不合法的数字
			return 0
		}
	}
	return result * i
}
