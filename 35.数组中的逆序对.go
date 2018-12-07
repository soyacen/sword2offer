package main

import "fmt"

/**
题目描述
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。
输入一个数组,求出这个数组中的逆序对的总数P。并将P对1000000007取模的结果输出。
即输出P%1000000007
输入描述:
题目保证输入的数组中没有的相同的数字

数据范围：

	对于%50的数据,size<=10^4

	对于%75的数据,size<=10^5

	对于%100的数据,size<=2*10^5

示例1
	输入
	1,2,3,4,5,6,7,0
	输出
	7
*/
func main() {
	fmt.Println(InversePairs([]int{1, 2, 3, 4, 5, 6, 7, 0}))
	fmt.Println(InversePairs([]int{4, 7, 5, 6, 4}))
	fmt.Println(InversePairs([]int{1, 3, 7, 8, 2, 4, 6, 5}))
	fmt.Println(InversePairs2([]int{1, 2, 3, 4, 5, 6, 7, 0}))
	fmt.Println(InversePairs2([]int{4, 7, 5, 6, 4}))
	fmt.Println(InversePairs2([]int{1, 3, 7, 8, 2, 4, 6, 5}))
}

// 顺序扫描整个数组 O(n2)
func InversePairs(data []int) int {
	P := 0
	q := 0
	length := len(data)
	for i := 0; i < length-1; i++ {
		q = 0
		for j := i + 1; j < length; j++ {
			if data[i] > data[j] {
				q++
			}
		}
		P += q
	}
	return P
}

// 分治思想，采用归并排序的思路来处理 O(nlogn)
func InversePairs2(data []int) int {
	length := len(data)
	if length < 2 {
		return 0
	}
	P := InversePairs(data[:length/2])
	P += InversePairs(data[length/2:])
	P += MergePairs(data[:length/2], data[length/2:])
	return P
}

func MergePairs(left, right []int) (result int) {
	for _, l := range left {
		for _, r := range right {
			if l > r {
				result++
			}
		}
	}
	return result
}
