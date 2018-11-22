package main

import (
	"fmt"
	"sync"
)

//用两个栈来实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型。

type Stack struct {
	Data []*StackItem
	Lock sync.RWMutex
}

type StackItem struct {
	num int
}

func (s *Stack) Pop() (result *StackItem) {
	if len(s.Data) == 0 {
		return
	}
	s.Lock.RLock()
	result = s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	s.Lock.RUnlock()
	return
}

func (s *Stack) Push(item *StackItem) {
	s.Lock.Lock()
	s.Data = append(s.Data, item)
	s.Lock.Unlock()
}

type Queue struct {
	stack1 *Stack
	stack2 *Stack
}

func (q *Queue) Enqueue(n int) {
	var tmp *StackItem
	for {
		tmp = q.stack1.Pop()
		if tmp == nil {
			break
		}
		q.stack2.Push(tmp)
	}
	q.stack1.Push(&StackItem{n})

	for {
		tmp = q.stack2.Pop()
		if tmp == nil {
			break
		}
		q.stack1.Push(tmp)
	}

}

func (q *Queue) Dequeue() (result int) {
	result = q.stack1.Pop().num
	return
}

func main() {
	s1 := &Stack{Lock: sync.RWMutex{}}
	s2 := &Stack{Lock: sync.RWMutex{}}
	q := &Queue{s1, s2}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(6)
	q.Enqueue(7)
	q.Enqueue(9)
	q.Enqueue(0)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
