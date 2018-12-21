package main

import (
	"fmt"
	"sort"
)

/**
题目描述
在一个长度为n的数组里的所有数字都在0到n-1的范围内。 数组中某些数字是重复的，但不知道有几个数字是重复的。
也不知道每个数字重复几次。请找出数组中任意一个重复的数字。 例如，如果输入长度为7的数组{2,3,1,0,2,5,3}，
那么对应的输出是第一个重复的数字2。
*/
func main() {
	arr := []int{2, 3, 1, 0, 2, 5, 3}
	dup := make([]int, 0, 1)
	fmt.Println(duplicate1(arr, &dup))
	fmt.Println(dup)

	dup = make([]int, 0, 1)
	fmt.Println(duplicate2(arr, &dup))
	fmt.Println(dup)

	dup = make([]int, 0, 1)
	fmt.Println(duplicate3(arr, &dup))
	fmt.Println(dup)

	dup = make([]int, 0, 1)
	fmt.Println(duplicate4(arr, &dup))
	fmt.Println(dup)
}

// 排序，相邻的相等的, 排序的时间复杂度为O(nlogn)
func duplicate1(numbers []int, duplication *[]int) bool {
	// write code here
	//这里要特别注意~找到任意重复的一个值并赋值到duplication[0]
	//函数返回True/False
	sort.Ints(numbers)
	for i := 1; i < len(numbers); i++ {
		if numbers[i] == numbers[i-1] {
			*duplication = append(*duplication, numbers[i-1])
			return true
		}
	}

	return false
}

// map Hash查询时间仅用O(1)，算法时间复杂度为O(n)，但是需要一个哈希表，空间复杂度为O(n)；
func duplicate2(numbers []int, duplication *[]int) bool {
	// write code here
	//这里要特别注意~找到任意重复的一个值并赋值到duplication[0]
	//函数返回True/False
	m := make(map[int]int)
	for _, v := range numbers {
		if _, ok := m[v]; ok {
			*duplication = append(*duplication, v)
			return true
		} else {
			m[v] = v
		}
	}
	return false
}

// 循环
func duplicate3(numbers []int, duplication *[]int) bool {
	for i, v := range numbers {
		for j, w := range numbers {
			if i != j && v == w {
				*duplication = append(*duplication, v)
				return true
			}
		}
	}
	return false
}

/**
利用元素数字都在0到n-1的范围内的特点，若不存在重复数字，则排序后的数字就在与其相同的索引值的位置，
即数字i在第i个位置。遍历的过程中寻找位置和元素不相同的值，并进行交换排序。时间复杂度O(n)，空间复杂度O(1)。


以数组｛2，3，1，0，2，5，3｝为例来分析找到重复数字的步骤。
数组的第0个数字（从0开始计数，和数组的下标保持一致）是2，与它的下标不想等，于是把它和下标为2的数字1交换，交换后的数组是｛1，3，2，0，2，5，3｝。
此时第0 个数字是1，仍然与它的下标不想等，继续把它和下标为1的数字3交换，得到数组｛0，1，2，3，2，5，3｝。
此时第0 个数字为0，接着扫描下一个数字，在接下来的几个数字中，下标为1，2，3的三个数字分别为1，2，3，他们的下标和数值都分别相等，因此不需要做任何操作。
接下来扫描下标为4的数字2.由于它的值与它的下标不登，再比较它和下标为2的数字。
注意到此时数组中下标为2的数字也是2，也就是数字2和下标为2和下标4的两个位置了，因此找到一个重复的数字。
*/
func duplicate4(numbers []int, duplication *[]int) bool {
	i := 0
	for i < len(numbers) {
		if i == numbers[i] {
			i++
		} else {
			if numbers[i] != numbers[numbers[i]] {
				numbers[i], numbers[numbers[i]] = numbers[numbers[i]], numbers[i]
			} else {
				*duplication = append(*duplication, numbers[i])
				return true
			}
		}
	}
	return false
}
