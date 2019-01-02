package main

import (
	. "github.com/yacen/sword2offer/list"
)

// 输入两个单调递增的链表，输出两个链表合成后的链表，当然我们需要合成后的链表满足单调不减规则。
func main() {
	head1 := &Node{10, &Node{12, &Node{31, &Node{45, &Node{58, nil}}}}}
	head2 := &Node{15, &Node{28, &Node{35, &Node{49, &Node{56, nil}}}}}

	//PrintListFromTailToHead(Merge(head1, head2))
	PrintListFromTailToHead(Merge2(head1, head2))
}

func Merge(pHead1 *Node, pHead2 *Node) (newList *Node) {
	node1 := pHead1
	node2 := pHead2
	newTail := newList
	for {
		var minNode *Node
		if node1.Data.(int) <= node2.Data.(int) {
			minNode = node1
			node1 = node1.Next
		} else {
			minNode = node2
			node2 = node2.Next
		}
		if newList == nil {
			newList = minNode
			newTail = newList
		} else {
			newTail.Next = minNode
			newTail = minNode
		}
		if node1 == nil {
			newTail.Next = node2
			break
		}
		if node2 == nil {
			newTail.Next = node1
			break
		}
	}
	return
}

func Merge2(pHead1 *Node, pHead2 *Node) (newList *Node) {
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	if pHead1.Data.(int) <= pHead2.Data.(int) {
		newList = pHead1
		newList.Next = Merge2(pHead1.Next, pHead2)
	} else {
		newList = pHead2
		newList.Next = Merge2(pHead1, pHead2.Next)
	}
	return
}
