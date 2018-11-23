package main

import "fmt"

// 大家都知道斐波那契数列，现在要求输入一个整数n，请你输出斐波那契数列的第n项（从0开始，第0项为0）。
// n<=39

func main() {
	for i := 0; i <= 39; i++ {
		fmt.Println(Fibonacci(i))
	}
}

func Fibonacci(n int) (result int64) {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
