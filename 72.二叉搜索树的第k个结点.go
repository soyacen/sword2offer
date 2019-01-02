package main

import (
	"fmt"
	"github.com/yacen/sword2offer/list"
	"github.com/yacen/sword2offer/tree"
)

/**
题目描述
给定一棵二叉搜索树，请找出其中的第k小的结点。例如， （5，3，7，2，4，6，8）
中，按结点数值大小顺序第三小结点的值为4。

				4
			2		6
			3		5	7
							8

*/
func main() {
	n5 := &tree.Node{Data: 5}
	n3 := &tree.Node{Data: 3}
	n7 := &tree.Node{Data: 7}
	n2 := &tree.Node{Data: 2}
	n4 := &tree.Node{Data: 4}
	n6 := &tree.Node{Data: 6}
	n8 := &tree.Node{Data: 8}

	n4.Left = n2
	n4.Right = n6

	n2.Right = n3

	n6.Left = n5
	n6.Right = n7

	n7.Right = n8

	fmt.Println(KthNode1(n4, 1).Data)
	fmt.Println(KthNode2(n4, 1).Data)
	fmt.Println(KthNode3(n4, 1).Data)

	fmt.Println(KthNode1(n4, 2).Data)
	fmt.Println(KthNode2(n4, 2).Data)
	fmt.Println(KthNode3(n4, 2).Data)

	fmt.Println(KthNode1(n4, 3).Data)
	fmt.Println(KthNode2(n4, 3).Data)
	fmt.Println(KthNode3(n4, 3).Data)

	fmt.Println(KthNode1(n4, 4).Data)
	fmt.Println(KthNode2(n4, 4).Data)
	fmt.Println(KthNode3(n4, 4).Data)

	fmt.Println(KthNode1(n4, 5).Data)
	fmt.Println(KthNode2(n4, 5).Data)
	fmt.Println(KthNode3(n4, 5).Data)

	fmt.Println(KthNode1(n4, 6).Data)
	fmt.Println(KthNode2(n4, 6).Data)
	fmt.Println(KthNode3(n4, 6).Data)

	fmt.Println(KthNode1(n4, 7).Data)
	fmt.Println(KthNode2(n4, 7).Data)
	fmt.Println(KthNode3(n4, 7).Data)
}

// 中序遍历整棵树
func KthNode1(pRoot *tree.Node, k int) *tree.Node {
	if k <= 0 {
		return nil
	}
	var result *tree.Node
	tree.InTree(pRoot, func(node *tree.Node) {
		k--
		if k == 0 {
			result = node
		}
	})

	return result
}

func KthNode2(root *tree.Node, k int) *tree.Node {
	if k <= 0 {
		return nil
	}
	stack := list.NewStack()

	node := root
	for node != nil {
		stack.Push(node)
		node = node.Left
	}
	var result *tree.Node
	for !stack.IsEmpty() && k > 0 {
		n := stack.Pop().(*tree.Node)
		k--
		if k == 0 {
			result = n
			break
		}
		if n.Right != nil {
			node := n.Right
			for node != nil {
				stack.Push(node)
				node = node.Left
			}
		}
	}
	return result
}

// 计算节点数
func KthNode3(root *tree.Node, k int) *tree.Node {
	count := countNode(root.Left)
	if k <= 0 {
		return nil
	}

	if count == k-1 {
		return root
	}

	if count >= k {
		return KthNode3(root.Left, k)
	} else {
		return KthNode3(root.Right, k-count-1)
	}

}

func countNode(root *tree.Node) int {
	i := 0
	tree.PreTree(root, func(node *tree.Node) {
		i++
	})
	return i
}
