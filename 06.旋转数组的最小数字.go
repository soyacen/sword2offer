package main

import "fmt"

// 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
// 输入一个非减排序的数组的一个旋转，输出旋转数组的最小元素。
// 例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。
// NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。

// 1。循环，2,最大值后面的，3。分治

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rotateArray := rotateArray(array, 7) //[]int{3,4,5,1,2}
	fmt.Println(minNumberInRotateArray(rotateArray))
}

func rotateArray(array []int, n int) (result []int) {
	result = append(result, array[n:]...)
	result = append(result, array[:n]...)
	return
}

func minNumberInRotateArray(array []int) (min int) {
	if len(array) == 0 {
		return
	}
	left := 0
	right := len(array) - 1
	if array[left] <= array[right] {
		return array[0]
	}
	if len(array) == 2 && array[left] > array[right] {
		return array[1]
	}
	middle := (left + right) / 2
	minLeft := minNumberInRotateArray(array[:middle])
	minRight := minNumberInRotateArray(array[middle+1:])
	min = array[middle]
	if minLeft != 0 && minLeft < min {
		min = minLeft
	}
	if minRight != 0 && minRight < min {
		min = minRight
	}
	return
}
