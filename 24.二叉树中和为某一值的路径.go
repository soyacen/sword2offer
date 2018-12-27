package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/tree"
)

// 输入一颗二叉树的跟节点和一个整数，打印出二叉树中结点值的和为输入整数的所有路径。
// 路径定义为从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。
// (注意: 在返回值的list中，数组长度大的数组靠前)

/*
		1
	2		3
4	15		6
7           8
*/

func main() {
	n1 := &Node{Data: 1}
	n2 := &Node{Data: 2}
	n3 := &Node{Data: 3}
	n4 := &Node{Data: 4}
	n15 := &Node{Data: 15}
	n6 := &Node{Data: 6}
	n7 := &Node{Data: 7}
	n8 := &Node{Data: 8}

	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n15
	n3.Left = n6
	n4.Left = n7
	n6.Left = n8

	fmt.Println(FindPath(n1, 18))

}

func FindPath(root *Node, expectNumber int) (result [][]int) {
	paths := Pre(root, 1)
	fmt.Println(paths)
	for _, path := range paths {
		if path[0] == expectNumber {
			result = append(result, path[1:])
			for i := len(result) - 1; i > 0; i-- {
				if len(result[i]) > len(result[i-1]) {
					result[i], result[i-1] = result[i-1], result[i]
				}
			}
		}
	}
	return
}

func Pre(root *Node, index int) (path [][]int) {
	if root == nil {
		return
	}
	num, _ := root.Data.(int)
	path = make([][]int, 0)
	if root.Left == nil && root.Right == nil {
		leafPath := make([]int, index+1)
		leafPath[0] = num
		leafPath[index] = num
		path = append(path, leafPath)
		return
	}
	if root.Left != nil {
		leftPath := Pre(root.Left, index+1)
		for _, itemPath := range leftPath {
			itemPath[0] = itemPath[0] + num
			itemPath[index] = num
			path = append(path, itemPath)
		}
	}
	if root.Right != nil {
		rightPath := Pre(root.Right, index+1)
		for _, itemPath := range rightPath {
			itemPath[0] = itemPath[0] + num
			itemPath[index] = num
			path = append(path, itemPath)
		}
	}

	return
}
