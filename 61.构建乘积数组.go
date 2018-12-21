package main

import "fmt"

/**
题目描述
给定一个数组A[0,1,...,n-1],请构建一个数组B[0,1,...,n-1],其中B中的元素B[i]=A[0]*A[1]*...*A[i-1]*A[i+1]*...*A[n-1]。不能使用除法。
https://www.nowcoder.com/questionTerminal/94a4d381a68b47b7a8bed86f2975db46
*/
func main() {
	fmt.Println(multiply([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(multiply2([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(multiply3([]int{1, 2, 3, 4, 5, 6, 7, 8}))
}

// B[0] = 1 * A1 * A2 * A3 * A4
// B[1] = A0 * 1 * A2 * A4 * A4
// B[2] = A0 * A1 * 1 * A3 * A4
// B[3] = A0 * A1 * A2 * 1 * A4
// B[4] = A0 * A1 * A2 * A3 * 1
func multiply(A []int) []int {
	length := len(A)
	B := make([]int, length)
	for i, _ := range A {
		B[i] = cumulation(A[:i]) * cumulation(A[i+1:])
	}
	return B
}

func cumulation(arr []int) (result int) {
	result = 1
	for _, v := range arr {
		result *= v
	}
	return
}

//// 分别乘左下三角形和右上三角形
func multiply2(A []int) []int {
	length := len(A)
	B := make([]int, length)
	leftTriangle := make([]int, length)
	rightTriangle := make([]int, length)
	for i := 0; i < length; i++ {
		j := length - i - 1
		if i == 0 {
			leftTriangle[i] = 1
			rightTriangle[j] = 1
			continue
		}
		leftTriangle[i] = A[i-1] * leftTriangle[i-1]
		rightTriangle[j] = A[j+1] * rightTriangle[j+1]
	}

	for i := 0; i < length; i++ {
		B[i] = rightTriangle[i] * leftTriangle[i]
	}
	return B
}

// 优化
func multiply3(A []int) []int {
	length := len(A)
	B := make([]int, length)

	// 左下三角形
	/*for i := 0; i < length; i++ {
		if i == 0 {
			B[i] = 1
			continue
		}
		B[i] = A[i-1] * B[i-1]
	}*/
	B[0] = 1
	for i := 1; i < length; i++ {
		B[i] = A[i-1] * B[i-1]
	}

	// 右上三角形
	tmp := 1
	/*for i := length - 1; i >= 0; i-- {
		if i < length-1 {
			tmp = tmp * A[i+1]
		}
		B[i] = B[i] * tmp
	}*/
	for i := length - 2; i >= 0; i-- {
		tmp *= A[i+1]
		B[i] *= tmp
	}
	return B
}
