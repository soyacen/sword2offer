package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
题目描述
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
例如，字符串"+100","5e2","-123","3.1416"和"-1E-16"都表示数值。
但是"12e","1a3.14","1.2.3","+-5"和"12e+4.3"都不是。
*/
func main() {
	fmt.Println(isNumeric1("+100"), isNumeric2("+100"))
	fmt.Println(isNumeric1("5e2"), isNumeric2("5e2"))
	fmt.Println(isNumeric1("-123"), isNumeric2("-123"))
	fmt.Println(isNumeric1("3.1416"), isNumeric2("3.1416"))
	fmt.Println(isNumeric1("1E-16"), isNumeric2("1E-16"))
	fmt.Println(isNumeric1("-.3421"), isNumeric2("-.3421"))

	fmt.Println(isNumeric1("12e"), isNumeric2("12e"))
	fmt.Println(isNumeric1("1a3.14"), isNumeric2("1a3.14"))
	fmt.Println(isNumeric1("1.2.3"), isNumeric2("1.2.3"))
	fmt.Println(isNumeric1("+-5"), isNumeric2("+-5"))
	fmt.Println(isNumeric1("12e+4.3"), isNumeric2("12e+4.3"))
}

func isNumeric1(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func isNumeric2(str string) bool {
	str = strings.TrimSpace(str)
	i := 0
	e := false
	point := false
	for ; i < len(str); i++ {
		c := str[i]
		if c == '.' {
			point = true
			break
		}
		if c == 'e' || c == 'E' {
			e = true
			break
		}
	}

	if e {
		return isInteger(str[:i]) && isE(str[i+1:])
	} else if point {
		return isInteger(str[:i]) && isDecimal(str[i+1:])
	}
	return isInteger(str)
}

func isInteger(str string) bool {
	if len(str) == 0 {
		return false
	}
	i := 0
	if str[0] == '+' || str[0] == '-' {
		i++
	}
	for ; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}

func isDecimal(str string) bool {
	for _, n := range str {
		if n >= '0' && n <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}

func isE(str string) bool {
	return isInteger(str)
}
