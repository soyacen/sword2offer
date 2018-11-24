package main

import "fmt"

//一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶总共有多少种跳法。
// 第一步可以跳1个，剩下n-1个台阶；可以跳2个，剩下n-2个台阶...可以跳n个，剩下n-n个
// 第二步基于第一步每种可能剩下的台阶数m，以跳1个，剩下m-1个台阶；可以跳2个，剩下m-2个台阶...可以跳m个，剩下m-m个
// ...
// 每种可能最后剩下0个台阶，也就是跳完了，计1。

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i, " = ", jumpFloorII(i))
	}
}

func jumpFloorII(number int) int {
	if number <= 0 {
		return 1
	}
	sum := 0
	n := 1
	for {
		sum += jumpFloorII(number - n)
		n++
		if number-n < 0 {
			break
		}
	}
	return sum

}
