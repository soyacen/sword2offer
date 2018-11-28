package util

import "fmt"

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
