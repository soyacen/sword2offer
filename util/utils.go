package util

import (
	"fmt"
	"sync"
)

type ListNode struct {
	Data interface{}
	Next *ListNode
}

type TreeNode struct {
	Data  interface{}
	Left  *TreeNode
	Right *TreeNode
}

func PrintListFromTailToHead(listNode *ListNode) {
	for current := listNode; current != nil; current = current.Next {
		fmt.Print(current.Data, ", ")
	}
	fmt.Println("")
}

func PrintPreTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	PrintPreTree(root.Left)
	PrintPreTree(root.Right)
}

func PrintInTree(root *TreeNode) {
	if root == nil {
		return
	}
	PrintInTree(root.Left)
	fmt.Println(root.Data)
	PrintInTree(root.Right)
}

func PrintPostTree(root *TreeNode) {
	if root == nil {
		return
	}
	PrintPostTree(root.Left)
	PrintPostTree(root.Right)
	fmt.Print(root.Data, ", ")
}

// ================================ stack =======================================================
type Stack struct {
	Data []interface{}
	Lock *sync.RWMutex
}

func NewStack() *Stack {
	return &Stack{Lock: &sync.RWMutex{}}
}

func (s *Stack) Pop() (result interface{}) {
	if len(s.Data) == 0 {
		return
	}
	if s.Lock == nil {
		s.Lock = &sync.RWMutex{}
	}
	s.Lock.RLock()
	result = s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	s.Lock.RUnlock()
	return
}

func (s *Stack) Push(item interface{}) {
	if item == nil {
		return
	}
	if s.Lock == nil {
		s.Lock = &sync.RWMutex{}
	}
	s.Lock.Lock()
	s.Data = append(s.Data, item)
	s.Lock.Unlock()
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

func (q *Queue) Enqueue(item interface{}) {
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

func (s *Queue) IsEmpty() bool {
	return s.Stack1.IsEmpty()
}
