package main

import (
	"fmt"
	"sort"
	"time"
)

// 题目描述
// 输入n个整数，找出其中最小的K个数。
// 例如输入4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4,。
func main() {
	var t int64
	var r []int
	t1 := time.Now().UnixNano()
	r = GetLeastNumbers_Solution1([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4)
	t = time.Now().UnixNano() - t1
	fmt.Println(r, t)

	t2 := time.Now().UnixNano()
	r = GetLeastNumbers_Solution2([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4)
	t = time.Now().UnixNano() - t2
	fmt.Println(r, t)

	t4 := time.Now().UnixNano()
	r = GetLeastNumbers_Solution4([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4)
	t = time.Now().UnixNano() - t4
	fmt.Println(r, t)
}

// 循环数组，向放入前4个元素，往后都和result的最大值比较，比最大值小才放入，
func GetLeastNumbers_Solution1(input []int, k int) (result []int) {
	length := len(input)
	if length <= k {
		result = make([]int, length)
		copy(result, input)
		return
	}
	result = make([]int, k)
	for i, v := range input {
		if i < k {
			result[i] = v
		} else {
			sort.Ints(result)
			if v < result[k-1] {
				result[k-1] = v
			}
		}
	}
	return
}

// 快排，取小四个
func GetLeastNumbers_Solution2(input []int, k int) (result []int) {
	length := len(input)
	result = make([]int, length)
	copy(result, input)
	if length <= k {
		return
	}
	sort.Ints(result)
	result = result[:4]
	return
}

// 冒泡最小的四个，如果k大于一半，冒泡最大 len - k
func GetLeastNumbers_Solution4(input []int, k int) (result []int) {
	length := len(input)
	result = make([]int, length)
	copy(result, input)
	if length <= k {
		return
	}

	bableMin := true
	if k > length/2 {
		bableMin = false
	}

	if bableMin {
		for i := 0; i < k; i++ {
			for j := 1; j < length; j++ {
				if result[j-1] < result[j] {
					result[j-1], result[j] = result[j], result[j-1]
				}
			}
		}
	} else {
		for i := 0; i < length-k; i++ {
			for j := 1; j < length; j++ {
				if result[j-1] > result[j] {
					result[j-1], result[j] = result[j], result[j-1]
				}
			}
		}
	}

	if bableMin {
		result = result[length-k:]
	} else {
		result = result[:k]
	}
	return
}
