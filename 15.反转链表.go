package main

import (
	. "github.com/yacen/sword2offer/list"
)

//输入一个链表，反转链表后，输出新链表的表头。
func main() {
	head := &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}}
	//PrintListFromTailToHead(ReverseList(head))
	PrintListFromTailToHead(ReverseList2(head))
}

func ReverseList(pHead *Node) *Node {

	var preNode *Node
	var currentNode *Node
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

func ReverseList2(pHead *Node) *Node {
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
