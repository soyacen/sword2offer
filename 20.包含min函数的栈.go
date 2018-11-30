package main

import (
	"errors"
	"fmt"
)

//定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的min函数（时间复杂度应为O（1））。

type MinStack struct {
	Data       []int
	OrderIndex []int
}

func (s *MinStack) Push(value int) {
	s.Data = append(s.Data, value)
	index := len(s.Data) - 1
	s.OrderIndex = append(s.OrderIndex, index)
	for i := index; i > 0; i-- {
		if s.Data[s.OrderIndex[i]] < s.Data[s.OrderIndex[i-1]] {
			s.OrderIndex[i], s.OrderIndex[i-1] = s.OrderIndex[i-1], s.OrderIndex[i]
		}
	}

}

func (s *MinStack) Pop() (value int, err error) {
	if len(s.Data) == 0 {
		return 0, errors.New("empty stack")
	}
	length := len(s.Data)
	value = s.Data[length-1]
	s.Data = s.Data[:length-1]
	isMove := false
	for i := 0; i < length-1; i++ {
		if s.OrderIndex[i] == length-1 {
			isMove = true
		}
		if isMove {
			s.OrderIndex[i], s.OrderIndex[i+1] = s.OrderIndex[i+1], s.OrderIndex[i]
		}
	}
	s.OrderIndex = s.OrderIndex[:length-1]
	return
}

func (s *MinStack) Top() (value int, err error) {
	if len(s.Data) == 0 {
		return 0, errors.New("empty stack")
	}
	length := len(s.Data)
	value = s.Data[length-1]
	return
}

func (s *MinStack) Min() (min int, err error) {
	if len(s.OrderIndex) == 0 {
		return 0, errors.New("empty stack")
	}
	return s.Data[s.OrderIndex[0]], err
}

func main() {
	stack := MinStack{}
	stack.Push(6)
	stack.Push(2)
	stack.Push(5)
	stack.Push(7)
	stack.Push(8)
	stack.Push(3)
	stack.Push(9)
	stack.Push(1)
	fmt.Println(stack.Min())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Min())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Min())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Min())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Min())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Min())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Min())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Min())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Min())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Min())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Min())
	fmt.Println(stack.Pop())

	fmt.Println("================")
}
