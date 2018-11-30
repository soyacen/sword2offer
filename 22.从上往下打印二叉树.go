package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/util"
	"sync"
)

// 从上往下打印出二叉树的每个节点，同层节点从左至右打印。

func main() {
	n1 := &TreeNode{Data: 1}
	n2 := &TreeNode{Data: 2}
	n3 := &TreeNode{Data: 3}
	n4 := &TreeNode{Data: 4}
	n5 := &TreeNode{Data: 5}
	n6 := &TreeNode{Data: 6}
	n7 := &TreeNode{Data: 7}
	n8 := &TreeNode{Data: 8}

	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n4.Left = n7
	n6.Left = n8

	PrintFromTopToBottom(n1)
}

func PrintFromTopToBottom(root *TreeNode) {
	s1 := &Stack{Lock: &sync.RWMutex{}}
	s2 := &Stack{Lock: &sync.RWMutex{}}
	queue := &Queue{s1, s2}
	queue.Enqueue(root)
	for !queue.IsEmpty() {
		node := queue.Dequeue()
		if node == nil {
			break
		}
		treeNode, _ := node.(*TreeNode)
		fmt.Println(treeNode.Data)
		if treeNode.Left != nil {
			queue.Enqueue(treeNode.Left)
		}
		if treeNode.Right != nil {
			queue.Enqueue(treeNode.Right)
		}
	}
}
