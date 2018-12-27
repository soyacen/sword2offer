package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/tree"
)

/**
题目描述
给定一个二叉树和其中的一个结点，请找出中序遍历顺序的下一个结点并且返回。
注意，树中的结点不仅包含左右子结点，同时包含指向父结点的指针。
*/

/**
			1
		2		3
	4	   5 		6
7                  8
                  9

*/

func main() {
	root := &LinkNode{Data: 1}
	n2 := &LinkNode{Data: 2}
	n3 := &LinkNode{Data: 3}
	n4 := &LinkNode{Data: 4}
	n5 := &LinkNode{Data: 5}
	n6 := &LinkNode{Data: 6}
	n7 := &LinkNode{Data: 7}
	n8 := &LinkNode{Data: 8}
	n9 := &LinkNode{Data: 9}

	root.Left = n2
	root.Right = n3

	n2.Next = root
	n2.Left = n4
	n2.Right = n5

	n3.Next = root
	n3.Right = n6

	n4.Next = n2
	n4.Left = n7

	n5.Next = n2

	n6.Next = n3
	n6.Left = n8

	n7.Next = n4

	n8.Next = n6
	n8.Left = n9

	n9.Next = n8

	root.InTree(func(node *LinkNode) {
		fmt.Println(node)
	})

	fmt.Println("======================")
	fmt.Println(GetNext(n5))
}

func GetNext(pNode *LinkNode) *LinkNode {
	if pNode.Right != nil {
		// 右子树不为空，就在右子树

		// 找右子树最左边的节点
		node := pNode.Right
		for node.Left != nil {
			node = node.Left
		}
		return node
		/*
			// 中序遍历整个子树，找第二个节点
			result := pNode
			pNode.InTree(func(node *LinkNode) {
				fmt.Println(node)
				if n == 2 {
					result = node
				}
				n++
			})
			return result*/
	} else {
		// 否则就在父节点,寻找 最终父节点 是 某节点的 左节点
		for pNode.Next != nil {
			// 如果正好是左节点，就直接返回父节点
			if pNode.Next.Left == pNode {
				return pNode.Next
			}
			// 如果是右节点，找父节点不是右节点的那个
			pNode = pNode.Next
		}
	}

	return nil
}
