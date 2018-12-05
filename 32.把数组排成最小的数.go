package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
题目描述
输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

例如输入数组{3，32，321}，则打印出这三个数字能排成的最小数字为321323。
*/
func main() {
	fmt.Println(PrintMinNumber([]int{3, 32, 321}))
	fmt.Println(PrintMinNumber2([]int{3, 321, 32}))
}

func PrintMinNumber(numbers []int) string {
	if len(numbers) == 0 {
		return ""
	}
	ns := spliceNumber("", numbers)
	sort.Strings(ns)
	fmt.Println(ns)
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

type IntArray []int

func (a IntArray) Len() int {
	return len(a)
}

func (a IntArray) Less(i, j int) bool {
	return strings.Compare(fmt.Sprintf("%d%d", a[i], a[j]), fmt.Sprintf("%d%d", a[j], a[i])) < 0
}

func (a IntArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func PrintMinNumber2(numbers []int) string {
	length := len(numbers)
	if length == 0 {
		return ""
	}
	array := IntArray(numbers)
	sort.Sort(array)
	s := ""
	for _, v := range array {
		s = fmt.Sprintf("%s%d", s, v)
	}
	return s
}
