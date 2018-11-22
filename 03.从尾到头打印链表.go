package main

import "fmt"

// 输入一个链表，按链表值从尾到头的顺序返回一个ArrayList。

type ListNode struct {
	Data interface{}
	next *ListNode
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	printListFromTailToHead(head)
}

func printListFromTailToHead(listNode *ListNode) {
	for current := listNode; current != nil; current = current.next {
		fmt.Println(current.Data)
	}
}
