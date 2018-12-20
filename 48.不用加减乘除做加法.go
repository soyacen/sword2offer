package main

import "fmt"

func main() {
	fmt.Println(Add1(10, 20))
	fmt.Println(Add2(10, 20))
}

/**
https://www.cnblogs.com/youxin/p/3295649.html
*/
func Add1(num1, num2 int) int {
	if num2 == 0 {
		return num1
	}
	sum := num1 ^ num2
	carry := (num1 & num2) << 1
	return Add1(sum, carry)
}

func Add2(num1, num2 int) int {
	sum := 0
	for num2 != 0 {
		sum = num1 ^ num2
		num2 = (num1 & num2) << 1
		num1 = sum
	}
	return sum
}
