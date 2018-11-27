package main

import "fmt"

// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前半部分，
// 所有的偶数位于数组的后半部分，并保证奇数和奇数，偶数和偶数之间的相对位置不变。

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	reOrderArray2(slice)
	fmt.Println(slice)
}

func reOrderArray(array []int) {
	var tmp []int = make([]int, len(array))
	copy(tmp, array)
	i := 0
	for _, v := range tmp {
		if v%2 == 1 {
			array[i] = v
			i++
		}
	}
	for _, v := range tmp {
		if v%2 == 0 {
			array[i] = v
			i++
		}
	}
}

// 冒泡(插入)思想，偶数排到后面(奇数插到前面)
func reOrderArray2(array []int) {
	start := -1
	end := -1
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			if start == -1 {
				start = i
				end = i
			} else if start == end {
				end = i
			}
		} else {
			if start != -1 {
				tmp := array[i]
				oddI := end
				for oddI >= start {
					array[oddI+1] = array[oddI]
					oddI--
				}
				array[start] = tmp
				start++
				end++
			}
		}

	}
}
