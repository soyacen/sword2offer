package main

import "fmt"

func main() {
	fmt.Println(Add1(10, 20))
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
