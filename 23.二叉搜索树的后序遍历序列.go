package main

//输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。
// 如果是则输出Yes,否则输出No。假设输入的数组的任意两个数字都互不相同。

/*
			4
		2		7
	1   3		6	8
				5
*/
import (
	"fmt"
)

func main() {
	fmt.Println(VerifySquenceOfBST([]int{1, 3, 2, 5, 6, 8, 7, 4}))
	fmt.Println(VerifySquenceOfBST([]int{4, 8, 6, 12, 16, 14, 10}))
	fmt.Println(VerifySquenceOfBST([]int{4, 6, 7, 5}))
	fmt.Println(VerifySquenceOfBST([]int{1, 2, 3, 4, 5}))
	fmt.Println(VerifySquenceOfBST([]int{5, 4, 3, 2, 1}))
	fmt.Println(VerifySquenceOfBST([]int{5}))
	fmt.Println(VerifySquenceOfBST([]int{7, 4, 6, 5}))
	fmt.Println(VerifySquenceOfBST([]int{4, 6, 12, 8, 16, 14, 10}))
	fmt.Println(VerifySquenceOfBST(nil))
}

func VerifySquenceOfBST(sequence []int) bool {
	if sequence == nil {
		return false
	}
	if len(sequence) == 0 {
		return true
	}
	length := len(sequence)
	root := sequence[length-1]

	i := 0
	small := 0
	for ; i < length-1; i++ {
		if sequence[i] > root {
			small = i
			break
		}
	}
	smallSequence := sequence[:small]
	big := length - 1
	for ; i < length-1; i++ {
		if sequence[i] < root {
			big = i
			break
		}
	}
	if big != length-1 {
		return false
	}
	bigSequence := sequence[small:big]
	return VerifySquenceOfBST(smallSequence) && VerifySquenceOfBST(bigSequence)
}
