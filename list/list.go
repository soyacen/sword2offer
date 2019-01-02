package list

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

func PrintListFromTailToHead(listNode *Node) {
	for current := listNode; current != nil; current = current.Next {
		fmt.Print(current.Data, ", ")
	}
	fmt.Println("")
}

// ================================ stack =======================================================
type Stack struct {
	Data []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Pop() (result interface{}) {
	if len(s.Data) == 0 {
		return
	}
	result = s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return
}

func (s *Stack) Push(item interface{}) {
	if item == nil {
		return
	}
	s.Data = append(s.Data, item)
}

func (s *Stack) Top() (result interface{}) {
	if len(s.Data) == 0 {
		return
	}
	result = s.Data[len(s.Data)-1]
	return
}

func (s *Stack) Length() int {
	return len(s.Data)
}

func (s *Stack) IsEmpty() bool {
	return len(s.Data) == 0
}

// ============================ queue ===================================

type Queue struct {
	Stack1 *Stack
	Stack2 *Stack
}

/*func (q *Queue) Enqueue(item interface{}) {
	var tmp interface{}
	for {
		tmp = q.Stack1.Pop()
		if tmp == nil {
			break
		}
		q.Stack2.Push(tmp)
	}
	q.Stack1.Push(item)

	for {
		tmp = q.Stack2.Pop()
		if tmp == nil {
			break
		}
		q.Stack1.Push(tmp)
	}

}

func (q *Queue) Dequeue() (result interface{}) {
	result = q.Stack1.Pop()
	return
}
*/

func NewQueue() (q *Queue) {
	return &Queue{Stack1: NewStack(), Stack2: NewStack()}
}

func (q *Queue) Enqueue(item interface{}) {
	q.Stack1.Push(item)
	/*	var tmp interface{}
		for {
			tmp = q.Stack1.Pop()
			if tmp == nil {
				break
			}
			q.Stack2.Push(tmp)
		}
		q.Stack1.Push(item)

		for {
			tmp = q.Stack2.Pop()
			if tmp == nil {
				break
			}
			q.Stack1.Push(tmp)
		}*/

}

func (q *Queue) Dequeue() (result interface{}) {
	if q.Stack2.IsEmpty() {
		for !q.Stack1.IsEmpty() {
			q.Stack2.Push(q.Stack1.Pop())
		}
	}
	result = q.Stack2.Pop()
	return
}

func (s *Queue) IsEmpty() bool {
	return s.Stack1.IsEmpty() && s.Stack2.IsEmpty()
}
