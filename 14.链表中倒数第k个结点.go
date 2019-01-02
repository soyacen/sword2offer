package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/list"
)

//输入一个链表，输出该链表中倒数第k个结点。

func main() {
	head := &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}}
	fmt.Println(FindKthToTail(head, 4).Data)
	fmt.Println(FindKthToTail2(head, 4).Data)
}

func FindKthToTail(pListHead *Node, k int) *Node {
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
func FindKthToTail2(pListHead *Node, k int) *Node {
	current := 1
	var preNode *Node
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
