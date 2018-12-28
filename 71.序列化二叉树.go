package main

import (
	"fmt"
	"github.com/yacen/sword2offer/tree"
	"strings"
)

/**
题目描述
请实现两个函数，分别用来序列化和反序列化二叉树

// 参考 04.重建二叉树
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

	str := Serialize(root)
	fmt.Println(str)
	fmt.Println("=======================")

	dsRoot := Deserialize(str)

	tree.PrintPreTree(dsRoot)
	fmt.Println("")
	tree.PrintPreTree(root)

	fmt.Println("=======================")

	tree.PrintInTree(dsRoot)
	fmt.Println("")
	tree.PrintInTree(root)
}

func Serialize(root *tree.Node) string {
	if root == nil {
		return ""
	}
	preStr := ""
	tree.PreTree(root, func(node *tree.Node) {
		preStr = fmt.Sprintf("%s,%d", preStr, node.Data)
	})
	preStr = strings.TrimLeft(preStr, ",")

	inStr := ""
	tree.InTree(root, func(node *tree.Node) {
		inStr = fmt.Sprintf("%s,%d", inStr, node.Data)
	})
	inStr = strings.TrimLeft(inStr, ",")

	return preStr + ",#" + inStr + ","
}

func Deserialize(str string) *tree.Node {
	var preArr, inArr []*tree.Node
	arr := []*tree.Node{}
	num := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ',' {
			arr = append(arr, &tree.Node{Data: num})
			num = 0
			continue
		}
		if str[i] == '#' {
			preArr = arr
			arr = []*tree.Node{}
			continue
		}
		if str[i] > '9' || str[i] < '0' {
			return nil
		}
		num = num*10 + int(str[i]-'0')
	}
	inArr = arr
	for _, v := range preArr {
		fmt.Printf("%p \n", v)
	}
	fmt.Print("#")
	for _, v := range inArr {
		fmt.Printf("%p \n", v)
	}
	fmt.Println()
	result := tree.ReConstructBinaryTree(preArr, inArr)
	return result
}
