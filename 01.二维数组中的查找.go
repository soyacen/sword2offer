package main

import "fmt"

// 在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
// 请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

func main() {
	arr1 := [5]int{4, 7, 9, 13, 24}
	arr2 := [5]int{6, 10, 15, 19, 29}
	arr3 := [5]int{14, 15, 16, 21, 33}
	arr4 := [5]int{19, 24, 29, 38, 44}
	array := [4][5]int{arr1, arr2, arr3, arr4}

	fmt.Println(find(array, 16))
}

func find(array [4][5]int, num int) (int, int, bool) {
	length := len(array)
	subLen := len(array[0])
	// 小于最小值，大于最大值 直接返回找不到
	if num < array[0][0] || num > array[length-1][subLen-1] {
		return -1, -1, false
	}
	row := 0
	col := 0
	startRow := false
	startCol := false
	// 剔除比 num 小的整条横向的与纵向的
	for {
		if array[row][subLen-1] < num {
			row++
		} else {
			startRow = true
		}
		if array[length-1][col] < num {
			col++
		} else {
			startCol = true
		}
		if startRow && startCol {
			break
		}
	}

	// 二分查找剩下的矩阵
	for {
		left := col
		right := subLen - 1
		var middle int
		for left <= right {
			middle = (right + left) / 2
			if array[row][middle] == num {
				return row, middle, true
			} else if array[row][middle] > num {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
		row++
		if row >= length {
			break
		}
	}
	return -1, -1, false
}
