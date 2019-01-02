package main

import (
	"fmt"
	"github.com/yacen/sword2offer/list"
)

// 输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否可能为该栈的弹出顺序。
// 假设压入栈的所有数字均不相等。例如序列1,2,3,4,5是某栈的压入顺序，
// 序列4,5,3,2,1是该压栈序列对应的一个弹出序列，
// 但4,3,5,1,2就不可能是该压栈序列的弹出序列。（注意：这两个序列的长度是相等的）
func main() {
	fmt.Println(IsPopOrder([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(IsPopOrder([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}

func IsPopOrder(pushV []int, popV []int) bool {
	if len(pushV) == 0 || len(popV) == 0 || len(pushV) != len(popV) {
		return false
	}
	s := list.Stack{}
	currentPop := 0
	for _, v := range pushV {
		s.Push(v)
		for !s.IsEmpty() {
			if s.Top() == popV[currentPop] {
				s.Pop()
				currentPop++
			} else {
				break
			}
		}
	}
	if currentPop == len(popV) {
		return true
	} else {
		return false
	}
}
