package main

import "fmt"

/**
题目描述
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
例如输入一个长度为9的数组{1,2,3,2,2,2,5,4,2}。由于数字2在数组中出现了5次，
超过数组长度的一半，因此输出2。如果不存在则输出0。
*/
func main() {
	array := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	fmt.Println(MoreThanHalfNum_Solution(array))
}

func MoreThanHalfNum_Solution(numbers []int) int {
	m := make(map[int]int)
	for _, num := range numbers {
		_, ok := m[num]
		if ok {
			m[num]++
		} else {
			m[num] = 1
		}
	}
	length := len(numbers) / 2
	for num, count := range m {
		if count > length {
			return num
		}
	}
	return 0
}
