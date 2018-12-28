package tree

import "fmt"

type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

func PrintPreTree(root *Node) {
	PreTree(root, func(node *Node) {
		fmt.Println(node.Data)
	})
}

func PreTree(root *Node, f func(node *Node)) {
	if root == nil {
		return
	}
	f(root)
	PreTree(root.Left, f)
	PreTree(root.Right, f)
}

func PrintInTree(root *Node) {
	InTree(root, func(node *Node) {
		fmt.Println(node.Data)
	})
}

func InTree(root *Node, f func(node *Node)) {
	if root == nil {
		return
	}
	InTree(root.Left, f)
	f(root)
	InTree(root.Right, f)
}

func SymmetricalInTree(root *Node, f func(node *Node)) {
	if root == nil {
		return
	}
	PrintPostTree(root.Right)
	f(root)
	PrintPostTree(root.Left)
}

func PrintPostTree(root *Node) {
	PostTree(root, func(node *Node) {
		fmt.Println(node.Data)
	})
}

func PostTree(root *Node, f func(node *Node)) {
	if root == nil {
		return
	}
	PostTree(root.Left, f)
	PostTree(root.Right, f)
	f(root)
}

func ReConstructBinaryTree(pre []*Node, in []*Node) (root *Node) {
	if len(pre) == 0 || len(in) == 0 {
		return nil
	}
	root = pre[0]
	rootInIndex := -1
	var leftPreNodes []*Node
	var leftInNodes []*Node
	var rightPreNodes []*Node
	var rightInNodes []*Node
	for i, node := range in {
		if node.Data == root.Data {
			rootInIndex = i
		} else {
			if rootInIndex == -1 {
				leftInNodes = append(leftInNodes, node)
			} else {
				rightInNodes = append(rightInNodes, node)
			}
		}
	}

	for _, preNode := range pre[1:] {
		isLeft := false
		for _, inNode := range leftInNodes {
			if preNode.Data == inNode.Data {
				isLeft = true
				break
			}
		}
		if isLeft {
			leftPreNodes = append(leftPreNodes, preNode)
		} else {
			rightPreNodes = append(rightPreNodes, preNode)
		}
	}

	root.Left = ReConstructBinaryTree(leftPreNodes, leftInNodes)
	root.Right = ReConstructBinaryTree(rightPreNodes, rightInNodes)
	return
}
