package main

import (
	"fmt"
	"math"
)

/**
题目描述
把只包含质因子2、3和5的数称作丑数（Ugly Number）。例如6、8都是丑数，但14不是，因为它包含质因子7。
习惯上我们把1当做是第一个丑数。求按从小到大的顺序的第N个丑数。
*/
func main() {
	fmt.Println(GetUglyNumber_Solution(140))
	fmt.Println(GetUglyNumber_Solution2(140))
}

func GetUglyNumber_Solution(index int) (result int) {
	count := 0
	number := 1
	for {
		if IsUglyNumber(number) {
			count++
		}
		if count == index {
			return number
		}
		number++
	}
}

func IsUglyNumber(number int) bool {
	for number%2 == 0 {
		number /= 2
	}

	for number%3 == 0 {
		number /= 3
	}

	for number%5 == 0 {
		number /= 5
	}

	if number == 1 {
		return true
	} else {
		return false
	}
}

/**
想办法从上一个丑数推断出下一个丑数，而不需要从1开始遍历再判断。从1开始的10个丑数分别为1，2，3，4，5，6，8，9，10，12。可以发现除了1以外，丑数都是由某个丑数*2或者*3或者*5得到的。如2是丑数1*2得到的，3是丑数1*3得到的，4是丑数1*4得到的，5是丑数1*5得到的，6是丑数2*3得到的……

具体算法步骤：

（1）从第一个丑数1开始，求出1*2=2 ，1*3=3 ，1*5 = 5；

（2）取上面乘积中大于1的最小值2，作为第二个丑数（丑数是个递增序列，所以第i+1个丑数一定比第i个丑数）

（3）求丑数2之前的丑数与2、3、5的乘积：1*2=2 ，1*3=3 ，1*5 = 5； 2*2 = 4； 2*3 = 6； 2*5 =10；

（4）取上面乘积中大于2的最小值3，作为第三个丑数

       ……

       ……

（i）取出丑数i之前的丑数分别与2、3、5的乘积

（i+1）取乘积中大于i的最小值作为丑数

（i+2）重复(i)(i+1)的步骤直到计数器等于N

这种思路的关键在于怎样确保数组里面的丑数是排好序的。我们假设数组中已经有若干个丑数，排好序后存在数组中。
我们把现有的最大丑数记做M。现在我们来生成下一个丑数，该丑数肯定是前面某一个丑数乘以2、3或者5的结果。
我们首先考虑把已有的每个丑数乘以2。在乘以2的时候，能得到若干个结果小于或等于M的。由于我们是按照顺序生成的，
小于或者等于M肯定已经在数组中了，我们不需再次考虑；我们还会得到若干个大于M的结果，但我们只需要第一个大于M的结果，
因为我们希望丑数是按从小到大顺序生成的，其他更大的结果我们以后再说。我们把得到的第一个乘以2后大于M的结果，记为M2。
同样我们把已有的每一个丑数乘以3和5，能得到第一个大于M的结果M3和M5。
那么下一个丑数应该是M2、M3和M5三个数的最小者。
---------------------
*/

/*
1,   2,3,5
1,2   \,4   3,6,  5,10
1,2,3  \,4,6,  \,6,9   5,10,15
1,2,3,4   \,\,6,8  \,6,9,12, 5,10,15,20
1,2,3,4,5



*/
func GetUglyNumber_Solution2(index int) (result int) {
	uglyNumbers := make([]int, index)
	uglyNumbers[0] = 1
	for i := 1; i < index; i++ {
		min2, min3, min5 := 0, 0, 0
		for _, v := range uglyNumbers {
			if 2*v > uglyNumbers[i-1] {
				min2 = 2 * v
				break
			}
		}
		for _, v := range uglyNumbers {
			if 3*v > uglyNumbers[i-1] {
				min3 = 3 * v
				break
			}
		}
		for _, v := range uglyNumbers {
			if 5*v > uglyNumbers[i-1] {
				min5 = 5 * v
				break
			}
		}
		uglyNumbers[i] = Min(min2, min3, min5)
	}
	return uglyNumbers[index-1]
}

func Min(x, y, z int) int {
	if x > y {
		if y > z {
			return z
		} else {
			return y
		}
	} else {
		if x > z {
			return z
		} else {
			return x
		}
	}
}

func IsPrime(number int) bool {
	if number <= 3 {
		// 对于小于4的质数只有2，3
		return number > 1
	}
	if number%2 == 0 {
		// 偶数的直接返回false
		return false
	}
	if 1 < number%6 && number%6 < 5 {
		// 质数还有一个特点，就是它总是等于 6x-1 或者 6x+1，其中 x 是大于等于1的自然数。
		return false
	}
	sqrt := int(math.Ceil(math.Sqrt(float64(number))))
	for i := 5; i <= sqrt; i += 6 {
		// 步长可以设为 6，然后每次只判断 6 两侧的数即可。
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}
	return true
}
