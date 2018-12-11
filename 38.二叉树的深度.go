package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/util"
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
	n4 := &TreeNode{Data: 4,
		Left: &TreeNode{Data: 2,
			Left: &TreeNode{Data: 1, Right: &TreeNode{Data: 4,
				Left: &TreeNode{Data: 2,
					Left: &TreeNode{Data: 1}, Right: &TreeNode{Data: 3}},
				Right: &TreeNode{Data: 7, Left: &TreeNode{Data: 6, Left: &TreeNode{Data: 5}},
					Right: &TreeNode{Data: 8}}}}, Right: &TreeNode{Data: 3}},
		Right: &TreeNode{Data: 7, Left: &TreeNode{Data: 6, Left: &TreeNode{Data: 5}},
			Right: &TreeNode{Data: 8, Left: &TreeNode{Data: 4,
				Left: &TreeNode{Data: 2,
					Left: &TreeNode{Data: 1}, Right: &TreeNode{Data: 3}},
				Right: &TreeNode{Data: 7, Left: &TreeNode{Data: 6, Left: &TreeNode{Data: 5}},
					Right: &TreeNode{Data: 8, Left: &TreeNode{Data: 4,
						Left: &TreeNode{Data: 2,
							Left: &TreeNode{Data: 1}, Right: &TreeNode{Data: 3}},
						Right: &TreeNode{Data: 7, Left: &TreeNode{Data: 6, Left: &TreeNode{Data: 5, Left: &TreeNode{Data: 4,
							Left: &TreeNode{Data: 2,
								Left: &TreeNode{Data: 1}, Right: &TreeNode{Data: 3}},
							Right: &TreeNode{Data: 7, Left: &TreeNode{Data: 6, Left: &TreeNode{Data: 5}},
								Right: &TreeNode{Data: 8}}}}},
							Right: &TreeNode{Data: 8}}}}}}}}}
	fmt.Println(TreeDepth(n4))
	fmt.Println(TreeDepth2(n4))
}

// 递归
func TreeDepth(pRoot *TreeNode) (result int) {
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
func TreeDepth2(pRoot *TreeNode) (result int) {
	if pRoot == nil {
		return
	}
	stack := []*TreeNode{pRoot}
	for len(stack) > 0 {
		result++
		tmp := []*TreeNode{}
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
