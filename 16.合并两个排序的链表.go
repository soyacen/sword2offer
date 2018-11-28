package main

import (
	. "github.com/yacen/sword2offer/util"
)

// 输入两个单调递增的链表，输出两个链表合成后的链表，当然我们需要合成后的链表满足单调不减规则。
func main() {
	head1 := &ListNode{10, &ListNode{12, &ListNode{31, &ListNode{45, &ListNode{58, nil}}}}}
	head2 := &ListNode{15, &ListNode{28, &ListNode{35, &ListNode{49, &ListNode{56, nil}}}}}

	PrintListFromTailToHead(Merge(head1, head2))
}

func Merge(pHead1 *ListNode, pHead2 *ListNode) (newList *ListNode) {
	node1 := pHead1
	node2 := pHead2
	newTail := newList
	for {
		var minNode *ListNode
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
