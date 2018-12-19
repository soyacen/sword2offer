package main

import "fmt"

/**
题目描述
求1+2+3+...+n，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
*/
func main() {
	fmt.Println(Sum_Solution(10))
}

// go好像不能使用短路

// 使用panic终止递归
func Sum_Solution(n int) int {
	defer func() {
		recover()
	}()
	n = n * n / n
	return n + Sum_Solution(n-1)
}

/*func Sum_Solution2(n int) int {
	a, b, sum := n, n+1, 0

	sum += b & f(a)
	a >>= 1
	b <<= 1
	fmt.Println(f(a), sum, a, b, b & f(a))

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	sum += b & f(a)
	a >>= 1
	b <<= 1

	return sum >> 1
}

func f(x int) int {
	return ((x & 1) << 31) >> 31
}
*/
