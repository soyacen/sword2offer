package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/util"
	"sync"
)

//用两个栈来实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型。

func main() {
	s1 := &Stack{Lock: &sync.RWMutex{}}
	s2 := &Stack{Lock: &sync.RWMutex{}}
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
