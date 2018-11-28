package main

import (
	. "github.com/yacen/sword2offer/util"
)

//输入一个链表，反转链表后，输出新链表的表头。
func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	//PrintListFromTailToHead(ReverseList(head))
	PrintListFromTailToHead(ReverseList2(head))
}

func ReverseList(pHead *ListNode) *ListNode {

	var preNode *ListNode
	var currentNode *ListNode
	restList := pHead
	for {
		currentNode = restList
		restList = restList.Next
		currentNode.Next = preNode
		preNode = currentNode
		if restList == nil {
			break
		}
	}
	return currentNode
}

func ReverseList2(pHead *ListNode) *ListNode {
	if pHead.Next == nil {
		return pHead
	} else {
		tail := pHead.Next
		node := ReverseList2(tail)
		tail.Next = pHead
		pHead.Next = nil
		return node
	}
}
