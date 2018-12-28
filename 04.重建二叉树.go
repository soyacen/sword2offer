package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/tree"
)

// 输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。
// 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
// 例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。

/*
		1
	2		3
4	5		6
7           8
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

	pre := []*Node{n1, n2, n4, n7, n3, n5, n6, n8}
	in := []*Node{n4, n7, n2, n1, n5, n3, n8, n6}
	root := ReConstructBinaryTree(pre, in)
	PrintPreTree(root)
	fmt.Println("==================")
	PrintInTree(root)
	//PrintInTree(root)
	/*	n1.left = n2
		n1.right=n3
		n2.left=n4
		n3.left=n5
		n3.right=n6
		n4.right=n7
		n6.left=n8*/
	//PrintPreTree(n1)
	//PrintInTree(n1)

	fmt.Println("==================")

	n1 = &Node{Data: 1}
	n2 = &Node{Data: 2}
	n22 := &Node{Data: 22}
	n3 = &Node{Data: 3}
	n33 := &Node{Data: 33}
	n4 = &Node{Data: 4}
	n44 := &Node{Data: 44}
	n5 = &Node{Data: 5}
	n55 := &Node{Data: 55}

	root = ReConstructBinaryTree([]*Node{n1, n2, n3, n5, n4, n22, n44, n33, n55}, []*Node{n5, n3, n2, n4, n1, n44, n22, n33, n55})
	PrintPreTree(root)
	fmt.Println("==================")
	PrintInTree(root)
}
