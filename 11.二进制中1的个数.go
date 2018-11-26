package main

import "fmt"

// 输入一个整数，输出该数二进制表示中1的个数。其中负数用补码表示。
func main() {
	for i := -10; i < 20; i++ {
		fmt.Println(i, "======>", NumberOf1(i))
	}
}

func NumberOf1(n int) (sum int) {

	for i := 0; i < 32; i++ {
		sum += (n & 1)
		n >>= 1
	}
	return
}
