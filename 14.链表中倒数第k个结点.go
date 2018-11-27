package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/util"
)

//输入一个链表，输出该链表中倒数第k个结点。

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Println(FindKthToTail(head, 4).Data)
	fmt.Println(FindKthToTail2(head, 4).Data)
}

func FindKthToTail(pListHead *ListNode, k int) *ListNode {
	length := 0
	for node := pListHead; node != nil; node = node.Next {
		length++
	}
	current := 0
	for node := pListHead; node != nil; node = node.Next {
		if current == length-4 {
			return node
		}
		current++
	}
	return nil
}

// 两个指针相差4个，一次遍历即可
func FindKthToTail2(pListHead *ListNode, k int) *ListNode {
	current := 1
	var preNode *ListNode
	for node := pListHead; node != nil; node = node.Next {
		if current == 4 {
			preNode = pListHead
		} else if current > 4 {
			preNode = preNode.Next
		}
		current++
	}
	return preNode
}
