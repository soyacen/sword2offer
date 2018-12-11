package main

import "fmt"

/**
题目描述
统计一个数字在排序数组中出现的次数。
*/
func main() {
	fmt.Println(GetNumberOfK([]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 7, 8, 9, 9}, 5))
	fmt.Println(GetNumberOfK([]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 7, 8, 9, 9}, 2))
	fmt.Println(GetNumberOfK([]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 7, 8, 9, 9}, 9))
	fmt.Println()

	fmt.Println(GetNumberOfK2([]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 7, 8, 9, 9}, 5))
	fmt.Println(GetNumberOfK2([]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 7, 8, 9, 9}, 2))
	fmt.Println(GetNumberOfK2([]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 7, 8, 9, 9}, 9))
	fmt.Println()

	fmt.Println(GetNumberOfK3([]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 7, 8, 9, 9}, 5))
	fmt.Println(GetNumberOfK3([]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 7, 8, 9, 9}, 2))
	fmt.Println(GetNumberOfK3([]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 7, 8, 9, 9}, 9))
	fmt.Println()
	fmt.Println(GetNumberOfK4([]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 7, 8, 9, 9}, 5))
	fmt.Println(GetNumberOfK4([]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 7, 8, 9, 9}, 2))
	fmt.Println(GetNumberOfK4([]int{1, 2, 3, 4, 5, 5, 5, 6, 6, 7, 8, 9, 9}, 9))
}

// 遍历1边数组
func GetNumberOfK(data []int, k int) (result int) {
	for _, v := range data {
		if k == v {
			result++
		}
	}
	return
}

// 2分查找
func GetNumberOfK2(data []int, k int) (result int) {
	middle, start, end := 0, -1, -1
	left := 0
	right := len(data) - 1

	for left < right {
		middle = (left + right) / 2
		if data[middle] == k {
			tmp := middle
			for tmp >= left && data[tmp] == k {
				start = tmp
				tmp--
			}
			tmp = middle
			for tmp <= right && data[tmp] == k {
				end = tmp
				tmp++
			}
			break
		} else if data[middle] > k {
			right = middle - 1
		} else if data[middle] < k {
			left = middle + 1
		}
	}
	return end - start + 1
}

// 首尾同时遍历
func GetNumberOfK3(data []int, k int) (result int) {
	start, end := -1, -1
	length := len(data)
	for i, v := range data {
		if k == v {
			if start == -1 {
				start = i
			}
		} else {
			if start != -1 {
				end = i - 1
				break
			}
		}
		if k == data[length-i-1] {
			if end == -1 {
				end = length - i - 1
			}
		} else {
			if end != -1 {
				start = length - i
				break
			}
		}
	}
	return end - start + 1
}

// 2分查找方案2
func GetNumberOfK4(data []int, k int) (result int) {
	middle, start, end := 0, -1, -1
	left := 0
	right := len(data) - 1

	for left < right {
		middle = (left + right) / 2
		if data[middle] == k {
			start, end = middle, middle
			if middle-1 >= left && data[middle-1] != k {
				start = middle
			} else {
				start = left + GetFirstK(data[left:middle], k)
			}

			if middle+1 <= right && data[middle+1] != k {
				end = middle
			} else {
				end = middle + 1 + GetLastK(data[middle+1:right], k)
			}
			break
		} else if data[middle] > k {
			right = middle - 1
		} else if data[middle] < k {
			left = middle + 1
		}
	}
	return end - start + 1
}

func GetFirstK(data []int, k int) (result int) {
	left := 0
	right := len(data) - 1
	middle := 0
	for left < right {
		middle = (left + right) / 2
		if data[middle] == k {
			if middle-1 > left {
				if data[middle-1] != k {
					return middle
				} else {
					right = middle - 1
				}
			} else {
				return middle
			}
		} else {
			left = middle + 1
		}
	}
	return
}

func GetLastK(data []int, k int) (result int) {
	left := 0
	right := len(data) - 1
	middle := 0
	for left < right {
		middle = (left + right) / 2
		if data[middle] == k {
			if middle+1 < right {
				if data[middle+1] != k {
					return middle
				} else {
					left = middle + 1
				}
			} else {
				return middle
			}
		} else {
			right = middle - 1
		}
	}
	return
}
