package main

import (
	"fmt"
	"github.com/yacen/sword2offer/tree"
)

/**
题目描述
从上到下按层打印二叉树，同一层结点从左至右输出。每一层输出一行。

				1
			2		22
		3	  4  44		33
	5						55

*/
func main() {
	root := &tree.Node{Data: 1}
	n2 := &tree.Node{Data: 2}
	n22 := &tree.Node{Data: 22}
	n3 := &tree.Node{Data: 3}
	n33 := &tree.Node{Data: 33}
	n4 := &tree.Node{Data: 4}
	n44 := &tree.Node{Data: 44}
	n5 := &tree.Node{Data: 5}
	n55 := &tree.Node{Data: 55}

	root.Left = n2
	root.Right = n22

	n2.Left = n3
	n2.Right = n4

	n22.Left = n44
	n22.Right = n33

	n3.Left = n5

	n33.Right = n55

	arr := PrintNormal(root)
	for _, a := range arr {
		for _, item := range a {
			fmt.Print(item.Data, ", ")
		}
		fmt.Println("")
	}

	fmt.Println("=================")

	arr1 := PrintNormal2(root)
	for _, a := range arr1 {
		for _, item := range a {
			fmt.Print(item.Data, ", ")
		}
		fmt.Println("")
	}
}

func PrintNormal(pRoot *tree.Node) [][]*tree.Node {
	if pRoot == nil {
		return nil
	}
	arr := []*tree.Node{pRoot}
	result := [][]*tree.Node{arr}
	left := false
	for {
		tmp := make([]*tree.Node, 0)
		for _, item := range arr {
			if item.Left != nil {
				tmp = append(tmp, item.Left)
			}
			if item.Right != nil {
				tmp = append(tmp, item.Right)
			}

		}
		if len(tmp) == 0 {
			break
		}
		result = append(result, tmp)
		arr = tmp
		left = !left
	}
	return result
}

func PrintNormal2(pRoot *tree.Node) [][]*tree.Node {
	if pRoot == nil {
		return nil
	}
	return PrintNormalItem([]*tree.Node{pRoot})
}

// num 奇数向左，偶数向右
func PrintNormalItem(nodes []*tree.Node) [][]*tree.Node {
	if len(nodes) == 0 {
		return nil
	}

	tmp := make([]*tree.Node, 0)
	for _, item := range nodes {
		if item.Left != nil {
			tmp = append(tmp, item.Left)
		}
		if item.Right != nil {
			tmp = append(tmp, item.Right)
		}

	}
	result := [][]*tree.Node{nodes}
	result = append(result, PrintNormalItem(tmp)...)
	return result
}
