package main

import (
	. "github.com/yacen/sword2offer/list"
)

// 输入一个链表，按链表值从尾到头的顺序返回一个ArrayList。

func main() {
	head := &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}}
	PrintListFromTailToHead(head)
}
