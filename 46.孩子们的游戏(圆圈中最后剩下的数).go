package main

import "fmt"

/**
题目描述
每年六一儿童节,牛客都会准备一些小礼物去看望孤儿院的小朋友,今年亦是如此。
HF作为牛客的资深元老,自然也准备了一些小游戏。其中,有个游戏是这样的:首先,让小朋友们围成一个大圈。
然后,他随机指定一个数m,让编号为0的小朋友开始报数。
每次喊到m-1的那个小朋友要出列唱首歌,然后可以在礼品箱中任意的挑选礼物,并且不再回到圈中,
从他的下一个小朋友开始,继续0...m-1报数....这样下去....直到剩下最后一个小朋友,
可以不用表演,并且拿到牛客名贵的“名侦探柯南”典藏版(名额有限哦!!^_^)。
请你试着想下,哪个小朋友会得到这份礼品呢？(注：小朋友的编号是从0到n-1)
*/
func main() {
	fmt.Println(LastRemaining_Solution(10, 4))
	fmt.Println(LastRemaining_Solution2(10, 4))
	fmt.Println(LastRemaining_Solution3(10, 4))
	fmt.Println(LastRemaining_Solution(10, 7))
	fmt.Println(LastRemaining_Solution2(10, 7))
	fmt.Println(LastRemaining_Solution3(10, 7))
}

// 递归数组
func LastRemaining_Solution(n, m int) int {
	children := make([]int, n)
	for i, _ := range children {
		children[i] = i
	}
	// 第一次从0开始
	return getChild(children, 0, m)
}

func getChild(children []int, start, m int) int {
	length := len(children)
	if length == 1 {
		return children[0]
	}
	// 当数到m-1时就出圈
	n := (m + start - 1) % length
	copy(children[n:], children[n+1:])
	// 以后都是从数到的数的下一个开始
	return getChild(children[:length-1], n, m)
}

/**
约瑟夫环

 * 记录一下其中映射关系的由来：
 * n个人排成一排，编号如下：
 * 0,1,2,3,4，.....,n-2,n-1
 * 按照规定，第一个出队的人的编号是 k = (m-1)%n,(对n取余是考虑到m>n的情况)，记最后剩下的人的编号为 F(n, m)
 * 接下来从编号为k+1的人开始报数，相当于如下的编号排列：
 * k+1, k+2, k+3, ...., n-2, n-1, 0, 1, ..., k-2, k-1  ////(*)
 * 对应的映射为：
 * 0, 1, 2, .....,n-k-3, n-k-2, n-k-1, n-k, ...., n-3, n-2  /////(**)
 * 描述这个映射关系的函数为：
 * p(x) = (x+n-k-1)%n, 这里的变量x表示上面数列（*）的元素，p(x) 表示上面数列（**）中相同位置对应的值
 * 然后我们实际上需要的映射是描述 (**)--->(*) 的函数，即为p^(x)
 * 如果找到这样的映射p^(x)，则F(n-1,m)可以通过p^(x)映射到(*)中，即跟F(n, m)建立联系。
 * 因为p(x) = (x+n-k-1)%n,根据取余运算的定义，它等价于 x+n-k-1 = g*n + p(x)，其中g表示某个整数
 * 等式左边移位化简后得：x = (g-1)*n + p(x) + k + 1
 * 注意x代表的是上面数列(*)中的编号，因此x的取值范围为[0,n-1],因此上面的等式需要对n取余，即：
 * x = [(g-1)*n + p(x) + k + 1]%n = [p(x) + k + 1]%n
 * 改写因变量和自变量得到：p^(x) = [x+k+1]%n
 * 上式中的x代表(**)数列中的元素，也即F(n-1,m),将 x = F(n-1, m) 和 k = (m-1)%n 代入即得：
 * F(n, m) = [ F(n-1, m) + (m-1)%n + 1 ]%n
 * 继续化简上式： 由 k=(m-1)%n 得 m-1=j*n+k,得 k = m-1-j*n = (m-1)%n，代入上式：
 * F(n, m) = [ F(n-1, m) + m-1-j*n + 1 ]%n = [ F(n-1, m) + m ]%n
 * 注意 [-j*n]%n = 0
 * 因此递归关系表达式为：F(n, m) = [F(n-1, m)+m]%n

要得到n个数字的序列的最后剩下的数字，只需要得到n-1个数字的序列的最后剩下的数字，并可以依此类推。
当n=1时，也就是序列中开始只有一个数字0，那么很显然最后剩下的数字就是0。我们把这种关系表示为：

         0                  n=1
f(n,m)={
         [f(n-1,m)+m]%n     n>1

这是一种时间复杂度为O(n)，空间复杂度为O(1)的方法，因此无论在时间上还是空间上都优于前面的思路。

*/
func LastRemaining_Solution2(n, m int) int {
	if n == 0 {
		return -1
	}
	s := 0
	for i := 2; i <= n; i++ {
		s = (s + m) % i
	}
	return s
}

func LastRemaining_Solution3(n, m int) int {
	if n == 1 {
		return 0
	}
	return (LastRemaining_Solution3(n-1, m) + m) % n
}
