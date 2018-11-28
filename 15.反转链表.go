package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/util"
)

//输入一个链表，反转链表后，输出新链表的表头。
func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	PrintListFromTailToHead2(ReverseList(head))
}

func ReverseList(pHead *ListNode) *ListNode {

	var preNode *ListNode
	var currentNode *ListNode
	restList := pHead
	for {
		currentNode = restList
		restList = restList.Next
		currentNode.Next = preNode
		PrintListFromTailToHead2(currentNode)
		preNode = currentNode
		if restList == nil {
			break
		}
	}
	return currentNode
}

func PrintListFromTailToHead2(listNode *ListNode) {
	for current := listNode; current != nil; current = current.Next {
		fmt.Print(current.Data, ", ")
	}
	fmt.Println("++++++++++++++++++++++++=")
}
