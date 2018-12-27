package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/tree"
)

/**
题目描述
输入一棵二叉树，求该树的深度。从根结点到叶结点依次经过的结点
（含根、叶结点）形成树的一条路径，最长路径的长度为树的深度。

*/

/*
			4
		2		7
	1   3		6	8
				5
*/

func main() {
	n4 := &Node{Data: 4,
		Left: &Node{Data: 2,
			Left: &Node{Data: 1, Right: &Node{Data: 4,
				Left: &Node{Data: 2,
					Left: &Node{Data: 1}, Right: &Node{Data: 3}},
				Right: &Node{Data: 7, Left: &Node{Data: 6, Left: &Node{Data: 5}},
					Right: &Node{Data: 8}}}}, Right: &Node{Data: 3}},
		Right: &Node{Data: 7, Left: &Node{Data: 6, Left: &Node{Data: 5}},
			Right: &Node{Data: 8, Left: &Node{Data: 4,
				Left: &Node{Data: 2,
					Left: &Node{Data: 1}, Right: &Node{Data: 3}},
				Right: &Node{Data: 7, Left: &Node{Data: 6, Left: &Node{Data: 5}},
					Right: &Node{Data: 8, Left: &Node{Data: 4,
						Left: &Node{Data: 2,
							Left: &Node{Data: 1}, Right: &Node{Data: 3}},
						Right: &Node{Data: 7, Left: &Node{Data: 6, Left: &Node{Data: 5, Left: &Node{Data: 4,
							Left: &Node{Data: 2,
								Left: &Node{Data: 1}, Right: &Node{Data: 3}},
							Right: &Node{Data: 7, Left: &Node{Data: 6, Left: &Node{Data: 5}},
								Right: &Node{Data: 8}}}}},
							Right: &Node{Data: 8}}}}}}}}}
	fmt.Println(TreeDepth(n4))
	fmt.Println(TreeDepth2(n4))
}

// 递归
func TreeDepth(pRoot *Node) (result int) {
	result = 0
	if pRoot == nil {
		return
	}
	result = 1
	l := TreeDepth(pRoot.Left)
	r := TreeDepth(pRoot.Right)
	if l > r {
		result += l
	} else {
		result += r
	}
	return
}

// 循环遍历
func TreeDepth2(pRoot *Node) (result int) {
	if pRoot == nil {
		return
	}
	stack := []*Node{pRoot}
	for len(stack) > 0 {
		result++
		tmp := []*Node{}
		for _, v := range stack {
			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}
		}

		stack = tmp
	}
	return
}
