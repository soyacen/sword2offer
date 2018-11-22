package main

import "fmt"

// 输入一个链表，按链表值从尾到头的顺序返回一个ArrayList。

type Node struct {
	Data interface{}
	next *Node
}

func main() {
	head := &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}}
	printListFromTailToHead(head)
}

func printListFromTailToHead(listNode *Node) {
	for current := listNode; current != nil; current = current.next {
		fmt.Println(current.Data)
	}
}
