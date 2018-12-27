package main

import (
	"fmt"
	"github.com/yacen/sword2offer/tree"
)

/**
题目描述
请实现一个函数，用来判断一颗二叉树是不是对称的。
注意，如果一个二叉树同此二叉树的镜像是同样的，定义其为对称的。

*/

/*
				1
			2		22
		3	  4  44		33
	5						55


*/

func main() {
	root := &tree.Node{Data: 1}
	n2 := &tree.Node{Data: 2}
	n22 := &tree.Node{Data: 2}
	n3 := &tree.Node{Data: 3}
	n33 := &tree.Node{Data: 3}
	n4 := &tree.Node{Data: 4}
	n44 := &tree.Node{Data: 4}
	n5 := &tree.Node{Data: 5}
	n55 := &tree.Node{Data: 5}

	root.Left = n2
	root.Right = n22

	n2.Left = n3
	n2.Right = n4

	n22.Left = n44
	n22.Right = n33

	n3.Left = n5

	n33.Right = n55

	fmt.Println(isSymmetrical1(root))
	fmt.Println(isSymmetrical2(root))

	n1 := &tree.Node{Data: 1}
	n2 = &tree.Node{Data: 2}
	n3 = &tree.Node{Data: 3}
	n4 = &tree.Node{Data: 4}
	n5 = &tree.Node{Data: 5}
	n6 := &tree.Node{Data: 6}
	n7 := &tree.Node{Data: 7}
	n8 := &tree.Node{Data: 8}

	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n4.Left = n7
	n6.Left = n8
	fmt.Println(isSymmetrical1(n1))
	fmt.Println(isSymmetrical2(n1))
}

// 从上往下遍历，放入数组，从数组的 首尾 向 中间遍历 ，不想等就返回false
func isSymmetrical1(pRoot *tree.Node) bool {

	nodes := []*tree.Node{pRoot.Left, pRoot.Right}
	length := len(nodes)
	for {
		tmpNodes := make([]*tree.Node, length*2)
		tmpLength := length * 2
		nilNum := 0
		for i := 0; i < length/2; i++ {
			left, right := nodes[i], nodes[length-1-i]
			if left == nil && right == nil {
				nilNum += 2
				continue
			} else if left == nil && right != nil {
				return false
			} else if left != nil && right == nil {
				return false
			} else {
				if left.Data == right.Data {
					tmpNodes[i*2] = left.Left
					tmpNodes[i*2+1] = left.Right

					tmpNodes[tmpLength-1-i*2] = right.Right
					tmpNodes[tmpLength-1-i*2-1] = right.Left
				} else {
					return false
				}
			}

		}
		if nilNum == length {
			break
		}
		length = tmpLength
		nodes = tmpNodes
	}

	return true
}

// 递归
func isSymmetrical2(pRoot *tree.Node) bool {
	if pRoot == nil {
		return true
	} else {
		return checkTree(pRoot.Left, pRoot.Right)
	}
}

func checkTree(leftTree *tree.Node, rightTree *tree.Node) bool {
	if leftTree == nil && rightTree == nil {
		return true
	} else if leftTree != nil && rightTree != nil && leftTree.Data == rightTree.Data {
		return checkTree(leftTree.Left, rightTree.Right) && checkTree(leftTree.Right, rightTree.Left)
	} else {
		return false
	}
}
