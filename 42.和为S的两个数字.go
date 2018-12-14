package main

import "fmt"

/**
题目描述
输入一个递增排序的数组和一个数字S，在数组中查找两个数，使得他们的和正好是S，
如果有多对数字的和等于S，输出两个数的乘积最小的。

输出描述:
对应每个测试案例，输出两个数，小的先输出。
*/
func main() {
	fmt.Println(FindNumbersWithSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11))
	fmt.Println(FindNumbersWithSum2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9))
}

func FindNumbersWithSum(array []int, sum int) (result []int) {
	for _, v := range array {
		n := sum - v
		if BinarySearch(array, n) != -1 {
			result = []int{v, n}
			break
		}
	}
	return
}

func BinarySearch(array []int, num int) (index int) {
	index = -1
	if len(array) == 0 {
		return
	}

	left := 0
	right := len(array) - 1
	var middle int
	for left <= right {
		middle = (left + right) / 2
		if array[middle] == num {
			index = middle
			return
		} else if array[middle] < num {
			left = middle + 1
		} else if array[middle] > num {
			right = middle - 1
		}
	}

	return
}

func FindNumbersWithSum2(array []int, sum int) (result []int) {
	first := 0
	second := len(array) - 1
	for first < second {
		s := array[first] + array[second]
		if s == sum {
			return []int{array[first], array[second]}
		} else if s < sum {
			first++
		} else {
			second--
		}
	}
	return
}
