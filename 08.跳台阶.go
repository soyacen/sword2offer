package main

import "fmt"

// 一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）。
// 第一步可以走1步，剩下n-1个台阶；可以走2步，剩下n-2个台阶。

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(jumpFloor(i))
		//fmt.Println(jumpFloor2(i))
	}
}

func jumpFloor(number int) (count int) {
	if number <= 0 {
		return 1
	}
	// 第一步跳1个
	count += jumpFloor(number - 1)
	if number > 1 {
		// 第一部跳2个
		count += jumpFloor(number - 2)
	}
	return count

}

func jumpFloor2(n int) int {
	if n <= 0 {
		return 0
	}
	f := 1
	g := 1
	for {
		n--
		g = g + f
		f = g - f
		if n <= 0 {
			break
		}
	}
	return f
}

// n g f
// 5 1 1
// 4 2 1
// 3 3 2
// 2 5 3
// 1 8 5
// 0 13 8
