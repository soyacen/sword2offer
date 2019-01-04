package main

import "fmt"

/**
题目描述
给定一个数组和滑动窗口的大小，找出所有滑动窗口里数值的最大值。

例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，

那么一共存在6个滑动窗口，他们的最大值分别为{4,4,6,6,6,5}；

针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下6个：
{[2,3,4],2,6,2,5,1}，
{2,[3,4,2],6,2,5,1}，
{2,3,[4,2,6],2,5,1}，
{2,3,4,[2,6,2],5,1}，
{2,3,4,2,[6,2,5],1}，
{2,3,4,2,6,[2,5,1]}。
*/
func main() {
	fmt.Println(maxInWindows1([]int{2, 3, 4, 2, 6, 2, 5, 1}, 3))
}

// 思路：
// go特有的slice最为每个滑动窗口
// 缓存每个滑动窗口的最大值下标
//
// 下一个滑动窗口新加的元素后，当上一个最大值不是上一轮滑动窗口最后一个元素 且 上一个最大值 大于 新加入的元素，则继续使用
// 否则就再找新的最大值，记录下标。
func maxInWindows1(num []int, size int) []int {
	length := len(num)
	if size > length {
		size = length
	}
	result := make([]int, 0, length-size+1)

	preMaxIndex := -1
	for i := 0; i <= length-size; i++ {
		if preMaxIndex >= 0 && preMaxIndex != i-1 && num[preMaxIndex] >= num[i+size-1] {
			result = append(result, num[preMaxIndex])
			continue
		}
		slice := num[i : i+size]
		maxIndex := 0
		for j, v := range slice {
			if v > slice[maxIndex] {
				maxIndex = j
			}
		}
		result = append(result, slice[maxIndex])
		preMaxIndex = maxIndex + i

	}
	return result
}
