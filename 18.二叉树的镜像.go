package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/util"
)

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

	PrintInTree(n1)
	fmt.Println("================")
	PrintInTree(Mirror2(n1))
}

func Mirror(pRoot *TreeNode) (result *TreeNode) {
	if pRoot == nil {
		return nil
	}
	pRoot.Left, pRoot.Right = pRoot.Right, pRoot.Left
	Mirror(pRoot.Left)
	Mirror(pRoot.Right)
	return pRoot
}

func Mirror2(pRoot *TreeNode) (result *TreeNode) {
	stack := NewStack()
	stack.Push(pRoot)
	for {
		node := stack.Pop()
		if node == nil {
			break
		}
		n, _ := node.(*TreeNode)
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
