package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/util"
)

/**
题目描述
输入一棵二叉树，判断该二叉树是否是平衡二叉树。


平衡二叉搜索树（Self-balancing binary search tree）又被称为
AVL树（有别于AVL算法），且具有以下性质：
它是一 棵空树或它的左右两个子树的高度差的绝对值不超过1，
并且左右两个子树都是一棵平衡二叉树。平
衡二叉树的常用实现方法有红黑树、AVL、替罪羊树、Treap、伸展树等。
最小二叉平衡树的节点总数的公式如下 F(n)=F(n-1)+F(n-2)+1
这个类似于一个递归的数列，可以参考Fibonacci(斐波那契)数列，
1是根节点，F(n-1)是左子树的节点数量，F(n-2)是右子树的节点数量。
*/
func main() {
	n4 := &TreeNode{Data: 4, Left: &TreeNode{Data: 2, Left: &TreeNode{Data: 1}, Right: &TreeNode{Data: 3}}, Right: &TreeNode{Data: 7, Left: &TreeNode{Data: 6, Left: &TreeNode{Data: 5}}, Right: &TreeNode{Data: 8}}}
	fmt.Println(IsBalanced_Solution(n4))

	n41 := &TreeNode{Data: 4,
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
							Right: &TreeNode{Data: 8, Left: &TreeNode{Data: 4, Left: &TreeNode{Data: 2, Left: &TreeNode{Data: 1}, Right: &TreeNode{Data: 3}}, Right: &TreeNode{Data: 7, Left: &TreeNode{Data: 6, Left: &TreeNode{Data: 5}}, Right: &TreeNode{Data: 8, Left: &TreeNode{Data: 4, Left: &TreeNode{Data: 2, Left: &TreeNode{Data: 1}, Right: &TreeNode{Data: 3}}, Right: &TreeNode{Data: 7, Left: &TreeNode{Data: 6, Left: &TreeNode{Data: 5}}, Right: &TreeNode{Data: 8}}}}}}}}}}}}}}}

	fmt.Println(IsBalanced_Solution(n41))
}

func IsBalanced_Solution(pRoot *TreeNode) (result bool) {
	if pRoot == nil {
		return true
	}
	//
	left := TreeDepth3(pRoot.Left)
	right := TreeDepth3(pRoot.Right)
	diff := left - right
	return (diff >= -1 && diff <= 1) && IsBalanced_Solution(pRoot.Left) && IsBalanced_Solution(pRoot.Right)
}

// 使用38题代码
func TreeDepth3(pRoot *TreeNode) (result int) {
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
