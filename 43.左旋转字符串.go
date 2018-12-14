package main

import "fmt"

/**
题目描述
汇编语言中有一种移位指令叫做循环左移（ROL），现在有个简单的任务，
就是用字符串模拟这个指令的运算结果。对于一个给定的字符序列S，
请你把其循环左移K位后的序列输出。例如，字符序列S=”abcXYZdef”,
要求输出循环左移3位后的结果，即“XYZdefabc”。是不是很简单？OK，搞定它！
*/
func main() {
	str := "abcdefghijklmn"
	fmt.Println(LeftRotateString(str, 6))
	fmt.Println(LeftRotateString2(str, 6))
	fmt.Println(LeftRotateString3(str, 6))
}

func LeftRotateString(str string, n int) (result string) {
	length := len(str)
	n = n % length
	return fmt.Sprintf("%s%s", str[n:], str[:n])
}

func LeftRotateString2(str string, n int) (result string) {
	runes := []rune(str)
	length := len(runes)
	n = n % length
	for n > 0 {
		n--
		tmp := runes[0]
		for i := 1; i < len(runes); i++ {
			runes[i-1] = runes[i]
		}
		runes[length-1] = tmp
	}
	return string(runes)
}

func LeftRotateString3(str string, n int) (result string) {
	runes := []rune(str)
	length := len(runes)
	n = n % length
	if n == 0 {
		return string(runes)
	}
	tmp := make([]rune, n)
	copy(tmp, runes)
	/*for i := n; i < length; i++ {
		runes[i-n] = runes[i]
	}*/
	copy(runes[:length-n], runes[n:])
	copy(runes[length-n:], tmp)
	return string(runes)
}
