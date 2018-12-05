package main

import (
	"fmt"
	"sort"
)

/*
题目描述
输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

例如输入数组{3，32，321}，则打印出这三个数字能排成的最小数字为321323。
*/
func main() {
	fmt.Println(PrintMinNumber([]int{3, 32, 321, 123}))
}

func PrintMinNumber(numbers []int) string {
	if len(numbers) == 0 {
		return ""
	}
	ns := spliceNumber("", numbers)
	sort.Strings(ns)
	return ns[0]
}

func spliceNumber(num string, numbers []int) (result []string) {
	length := len(numbers)
	if length == 0 {
		return []string{num}
	}
	for i, v := range numbers {
		n := fmt.Sprintf("%s%d", num, v)
		ns := remove(i, numbers)
		result = append(result, spliceNumber(n, ns)...)
	}
	return result
}

func remove(i int, numbers []int) (result []int) {
	result = append(result, numbers[:i]...)
	result = append(result, numbers[i+1:]...)
	return
}
