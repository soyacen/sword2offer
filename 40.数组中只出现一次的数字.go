package main

import "fmt"

/**
题目描述
一个整型数组里除了两个数字之外，其他的数字都出现了偶数次。请写程序找出这两个只出现一次的数字。

例如：[2, 4, 3, 6, 3, 2, 5, 5]->[4, 6]
*/
func main() {
	n1, n2 := FindNumsAppearOnce([]int{2, 4, 3, 6, 3, 2, 5, 5})
	fmt.Println(n1, n2)

	n1, n2 = FindNumsAppearOnce2([]int{2, 4, 3, 6, 3, 2, 5, 5})
	fmt.Println(n1, n2)
}

func FindNumsAppearOnce(data []int) (num1, num2 int) {
	for i, v1 := range data {
		have := false
		for j, v2 := range data {
			if i != j {
				if v1 == v2 {
					have = true
					break
				}
			}
		}
		if !have {
			if num1 == 0 {
				num1 = v1
			} else if num2 == 0 {
				num2 = v1
				return
			}
		}
	}
	return
}

/**
我们知道两个相同的数字异或的结果为0，所以如果只有一个出现一次的数字，就可以让所有的数字进行异或，那么最后得到的数字就是只出现一次的数字。

现在问题变成了有两个数字，同样我们可以让所有数字异或，最终的结果是两个只出现一次的数字的异或的结果。
我们可以找到这个结果的二进制位上为1的下标，例如异或的结果为4，那么对应到二进制上是100，所以在第三位上着两个数字不相同，一个为1，一个为0，
所以我们可以以二进制中第三位为区分标准，把数组划分成两份，二进制的第三位为1和为0。

当得到两个小组后，我们知道这两个只出现一次的数已经被分开了，这样就回到最原始的问题上了，问题迎刃而解。代码如下。

直接循环数组，异或每个元素
0    [2,    4,    3,    6,    3,    2,    5,    5]
0000 [0010, 0100, 0011, 0110, 0011, 0010, 0111, 0111]
      0010
			0110
                  0101
                        0011
                              0000
                                    0010
                                          0101
                                                0010
0010
1  0010 & 0001  => 0000
2  0001 & 0001  => 0001
第二位 0 [0100,]   ->   4
第二位 1 [0010, 0011, 0110, 0011, 0010, 0111, 0111]  ->
         0010  0001  0111  0100  0110  0001  0110  > 6

*/
func FindNumsAppearOnce2(data []int) (num1, num2 int) {
	if len(data) < 2 {
		return
	}
	bit := 0x0000
	for _, v := range data {
		bit ^= v
	}

	var firstOne uint
	for ; bit&0x0001 == 0; bit >>= 1 {
		firstOne++
	}
	fmt.Println(firstOne)
	for _, v := range data {
		ax := (v >> firstOne) & 0x0001
		if ax == 1 {
			num2 ^= v
		} else {
			num1 ^= v
		}
	}
	return
}
