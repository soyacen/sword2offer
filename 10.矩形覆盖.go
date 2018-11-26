package main

import "fmt"

//
// 我们可以用2*1的小矩形横着或者竖着去覆盖更大的矩形。请问用n个2*1的小矩形无重叠地覆盖一个2*n的大矩形，总共有多少种方法？

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(rectCover(i))
		//fmt.Println(rectCover2(i))
		fmt.Println(rectCover3(i))
	}
}

func rectCover(number int) int {
	if number <= 0 {
		return 1
	}
	sum := rectCover(number - 1)
	if number > 1 {
		sum += rectCover(number - 2)
	}
	return sum
}

func rectCover2(number int) int {
	if number <= 0 {
		return 0
	}
	if number <= 2 {
		return number
	}

	return rectCover2(number-1) + rectCover2(number-2)
}

func rectCover3(number int) (sum int) {
	if number <= 0 {
		return 0
	}
	prePre := 1
	if number == 1 {
		sum = 1
		return
	}
	pre := 2
	if number == 2 {
		sum = 2
		return
	}
	for n := 3; n <= number; n++ {
		sum = prePre + pre
		fmt.Println("===>", sum)
		prePre = pre
		pre = sum
	}
	return
}
