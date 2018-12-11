package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/util"
)

// 输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。
// 要求不能创建任何新的结点，只能调整树中结点指针的指向。

/*
			4
		2		7
	1   3		6	8
				5
*/

func main() {
	n4 := &TreeNode{Data: 4, Left: &TreeNode{Data: 2, Left: &TreeNode{Data: 1}, Right: &TreeNode{Data: 3}}, Right: &TreeNode{Data: 7, Left: &TreeNode{Data: 6, Left: &TreeNode{Data: 5}}, Right: &TreeNode{Data: 8}}}
	list := Convert(n4)
	for node := list; node != nil; node = node.Right {
		fmt.Println(node.Data)
		if node.Data == 8 {
			for n := node; n != nil; n = n.Left {
				fmt.Println(n.Data)
			}
		}
	}
}

func Convert(pRootOfTree *TreeNode) (result *TreeNode) {
	if pRootOfTree == nil {
		return
	}
	if pRootOfTree.Left == nil && pRootOfTree.Right == nil {
		return pRootOfTree
	}
	small := Convert(pRootOfTree.Left)
	if small != nil {
		node := small
		for {
			if node.Right == nil {
				node.Right = pRootOfTree
				pRootOfTree.Left = node
				break
			}
			node = node.Right
		}
	}
	big := Convert(pRootOfTree.Right)

	pRootOfTree.Right = big
	if big != nil {
		big.Left = pRootOfTree
	}
	/*


		if pRootOfTree.Left != nil {
			fmt.Print(pRootOfTree.Left.Data)
		}
		fmt.Print(" <- ",pRootOfTree.Data," -> ")

		if pRootOfTree.Right != nil {
			fmt.Print(pRootOfTree.Right.Data)
		}
		fmt.Println()*/
	return small
}
