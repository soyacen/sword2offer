package main

import "fmt"

// 题目描述
// HZ偶尔会拿些专业问题来忽悠那些非计算机专业的同学。
// 今天测试组开完会后,他又发话了:在古老的一维模式识别中,常常需要计算连续子向量的最大和,当向量全为正数的时候,问题很好解决。
// 但是,如果向量中包含负数,是否应该包含某个负数,并期望旁边的正数会弥补它呢？
// 例如:{6,-3,-2,7,-15,1,2,2},连续子向量的最大和为8(从第0个开始,到第3个为止)。
// 给一个数组，返回它的最大连续子序列的和，你会不会被他忽悠住？(子向量的长度至少是1)
func main() {
	fmt.Println(FindGreatestSumOfSubArray([]int{6, -3, -2, 7, -15, 1, 2, 2}))
	fmt.Println(FindGreatestSumOfSubArray2([]int{6, -3, -2, 7, -15, 1, 2, 2}))
	fmt.Println(FindGreatestSumOfSubArray([]int{1, -2, 3, 10, -4, 7, 2, -5}))
	fmt.Println(FindGreatestSumOfSubArray2([]int{1, -2, 3, 10, -4, 7, 2, -5}))
	fmt.Println(FindGreatestSumOfSubArray2([]int{-11, -2, -13, -110, -4, -1, -17, -12, -5}))
}

// 递归全部的可能
func FindGreatestSumOfSubArray(array []int) int {
	if len(array) == 1 {
		return array[0]
	}
	max := array[0]
	sum := 0
	for _, v := range array {
		sum += v
		if sum > max {
			max = sum
		}
	}
	subMax := FindGreatestSumOfSubArray(array[1:])
	if max > subMax {
		return max
	} else {
		return subMax
	}
}

// 循环数组，当和小于等于0时，就把和重置成当前元素，否则就加上当前元素，判断当前和是否大于上次记录的和，大于就赋值。
func FindGreatestSumOfSubArray2(array []int) int {
	if len(array) == 1 {
		return array[0]
	}
	sum := 0
	preMax := array[0]
	for _, v := range array {
		if sum <= 0 {
			sum = v
		} else {
			sum += v
		}
		if preMax < sum {
			preMax = sum
		}
	}
	return preMax
}
