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

// ================================ stack =======================================================
type Stack struct {
	Data []interface{}
	Lock *sync.RWMutex
}

func (s *Stack) Pop() (result interface{}) {
	if len(s.Data) == 0 {
		return
	}
	s.Lock.RLock()
	result = s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	s.Lock.RUnlock()
	return
}

func (s *Stack) Push(item interface{}) {
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

func (s *Stack) IsEmpty() bool {
	return len(s.Data) == 0
}
