package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/util"
)

// 输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。
// 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
// 例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。

/*    1
 2   3
4    5 6
 7     8*/

func main() {
	n1 := &TreeNode{Data: 1}
	n2 := &TreeNode{Data: 2}
	n3 := &TreeNode{Data: 3}
	n4 := &TreeNode{Data: 4}
	n5 := &TreeNode{Data: 5}
	n6 := &TreeNode{Data: 6}
	n7 := &TreeNode{Data: 7}
	n8 := &TreeNode{Data: 8}

	pre := []*TreeNode{n1, n2, n4, n7, n3, n5, n6, n8}
	in := []*TreeNode{n4, n7, n2, n1, n5, n3, n8, n6}
	root := reConstructBinaryTree(pre, in)
	printPreTree(root)
	fmt.Println("==================")
	printInTree(root)
	//printInTree(root)
	/*	n1.left = n2
		n1.right=n3
		n2.left=n4
		n3.left=n5
		n3.right=n6
		n4.right=n7
		n6.left=n8*/
	//printPreTree(n1)
	//printInTree(n1)
}

func printPreTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	printPreTree(root.Left)
	printPreTree(root.Right)
}

func printInTree(root *TreeNode) {
	if root == nil {
		return
	}
	printInTree(root.Left)
	fmt.Println(root.Data)
	printInTree(root.Right)
}

func reConstructBinaryTree(pre []*TreeNode, in []*TreeNode) (root *TreeNode) {
	if len(pre) == 0 || len(in) == 0 {
		return nil
	}
	root = pre[0]
	rootInIndex := -1
	var leftPreTreeNodes []*TreeNode
	var leftInTreeNodes []*TreeNode
	var rightPreTreeNodes []*TreeNode
	var rightInTreeNodes []*TreeNode
	for i, node := range in {
		if node == root {
			rootInIndex = i
		} else {
			if rootInIndex == -1 {
				leftInTreeNodes = append(leftInTreeNodes, node)
			} else {
				rightInTreeNodes = append(rightInTreeNodes, node)
			}
		}
	}

	for _, preNode := range pre[1:] {
		isLeft := false
		for _, inNode := range leftInTreeNodes {
			if preNode == inNode {
				isLeft = true
				break
			}
		}
		if isLeft {
			leftPreTreeNodes = append(leftPreTreeNodes, preNode)
		} else {
			rightPreTreeNodes = append(rightPreTreeNodes, preNode)
		}
	}

	root.Left = reConstructBinaryTree(leftPreTreeNodes, leftInTreeNodes)
	root.Right = reConstructBinaryTree(rightPreTreeNodes, rightInTreeNodes)
	return
}
