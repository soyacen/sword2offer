package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/list"
	. "github.com/yacen/sword2offer/tree"
)

/**
题目描述
操作给定的二叉树，将其变换为源二叉树的镜像。
输入描述:
二叉树的镜像定义：源二叉树
    	    8
    	   /  \
    	  6   10
    	 / \  / \
    	5  7 9 11
    	镜像二叉树
    	    8
    	   /  \
    	  10   6
    	 / \  / \
    	11 9 7  5
*/

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

	PrintInTree(n1)
	fmt.Println("================")
	PrintInTree(Mirror2(n1))
}

func Mirror(pRoot *Node) (result *Node) {
	if pRoot == nil {
		return nil
	}
	pRoot.Left, pRoot.Right = pRoot.Right, pRoot.Left
	Mirror(pRoot.Left)
	Mirror(pRoot.Right)
	return pRoot
}

func Mirror2(pRoot *Node) (result *Node) {
	stack := NewStack()
	stack.Push(pRoot)
	for {
		node := stack.Pop()
		if node == nil {
			break
		}
		n, _ := node.(*Node)
		n.Left, n.Right = n.Right, n.Left
		if n.Left != nil {
			stack.Push(n.Left)
		}
		if n.Right != nil {
			stack.Push(n.Right)
		}
	}

	return pRoot
}
