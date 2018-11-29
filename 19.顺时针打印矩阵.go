package main

import (
	"fmt"
)

// 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字，
// 例如，如果输入如下4 X 4矩阵：
// 1 2 3 4
// 5 6 7 8
// 9 10 11 12
// 13 14 15 16
// 则依次打印出数字
// 1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10.

func main() {
	matrix := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
		[]int{17, 18, 19, 20},
	}
	fmt.Printf("%v", printMatrix(matrix))
}

func printMatrix(matrix [][]int) (result []int) {
	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	for {
		i := top
		// 上面
		for j := left; j <= right; j++ {
			result = append(result, matrix[i][j])
		}
		i++
		// 右面
		for j := right; i < bottom; i++ {
			result = append(result, matrix[i][j])
		}
		// 下面
		for j := right; j >= left; j-- {
			result = append(result, matrix[i][j])
		}
		i = bottom - 1
		// 左面
		for j := left; i > top; i-- {
			result = append(result, matrix[i][j])
		}

		top++
		bottom--
		left++
		right--
		if top > bottom || left > right {
			break
		}
	}

	return
}
