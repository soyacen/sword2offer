package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/tree"
	. "github.com/yacen/sword2offer/util"
)

// 从上往下打印出二叉树的每个节点，同层节点从左至右打印。

func main() {
	n1 := &Node{Data: 1}
	n2 := &Node{Data: 2}
	n3 := &Node{Data: 3}
	n4 := &Node{Data: 4}
	n5 := &Node{Data: 5}
	n6 := &Node{Data: 6}
	n7 := &Node{Data: 7}
	n8 := &Node{Data: 8}

	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n4.Left = n7
	n6.Left = n8

	PrintFromTopToBottom(n1)
}

func PrintFromTopToBottom(root *Node) {
	queue := NewQueue()
	queue.Enqueue(root)
	for !queue.IsEmpty() {
		node := queue.Dequeue()
		if node == nil {
			break
		}
		Node, _ := node.(*Node)
		fmt.Println(Node.Data)
		if Node.Left != nil {
			queue.Enqueue(Node.Left)
		}
		if Node.Right != nil {
			queue.Enqueue(Node.Right)
		}
	}
}
