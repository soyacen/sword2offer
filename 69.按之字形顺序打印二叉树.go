package main

import (
	"fmt"
	"github.com/yacen/sword2offer/tree"
)

/**
题目描述
请实现一个函数按照之字形打印二叉树，即第一行按照从左到右的顺序打印，
第二层按照从右至左的顺序打印，第三行按照从左到右的顺序打印，其他行以此类推。

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

	arr := Print(root)
	for _, a := range arr {
		for _, item := range a {
			fmt.Print(item.Data, ", ")
		}
		fmt.Println("")
	}

	fmt.Println("=================")

	arr1 := Print2(root)
	for _, a := range arr1 {
		for _, item := range a {
			fmt.Print(item.Data, ", ")
		}
		fmt.Println("")
	}
}

func Print(pRoot *tree.Node) [][]*tree.Node {
	if pRoot == nil {
		return nil
	}
	arr := []*tree.Node{pRoot}
	result := [][]*tree.Node{arr}
	left := false
	for {
		tmp := make([]*tree.Node, 0)
		var item *tree.Node
		for i, _ := range arr {
			item = arr[len(arr)-1-i]
			if left {
				if item.Left != nil {
					tmp = append(tmp, item.Left)
				}
				if item.Right != nil {
					tmp = append(tmp, item.Right)
				}
			} else {
				if item.Right != nil {
					tmp = append(tmp, item.Right)
				}
				if item.Left != nil {
					tmp = append(tmp, item.Left)
				}
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

func Print2(pRoot *tree.Node) [][]*tree.Node {
	if pRoot == nil {
		return nil
	}
	return PrintItem([]*tree.Node{pRoot}, 1)
}

// num 奇数向左，偶数向右
func PrintItem(nodes []*tree.Node, num int) [][]*tree.Node {
	if len(nodes) == 0 {
		return nil
	}

	tmp := make([]*tree.Node, 0)
	for i, _ := range nodes {
		item := nodes[len(nodes)-1-i]

		if num%2 == 1 {
			if item.Right != nil {
				tmp = append(tmp, item.Right)
			}
			if item.Left != nil {
				tmp = append(tmp, item.Left)
			}
		} else {
			if item.Left != nil {
				tmp = append(tmp, item.Left)
			}
			if item.Right != nil {
				tmp = append(tmp, item.Right)
			}

		}
	}
	num++
	result := [][]*tree.Node{nodes}
	result = append(result, PrintItem(tmp, num)...)
	return result
}
