package tree

import "fmt"

type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

func PrintPreTree(root *Node) {
	PreTree(root, func(node *Node) {
		fmt.Println(node)
	})
}

func PreTree(root *Node, f func(node *Node)) {
	if root == nil {
		return
	}
	f(root)
	PrintPostTree(root.Left)
	PrintPostTree(root.Right)
}

func PrintInTree(root *Node) {
	InTree(root, func(node *Node) {
		fmt.Println(node)
	})
}

func InTree(root *Node, f func(node *Node)) {
	if root == nil {
		return
	}
	PrintPostTree(root.Left)
	f(root)
	PrintPostTree(root.Right)
}

func PrintPostTree(root *Node) {
	PostTree(root, func(node *Node) {
		fmt.Print(node.Data, ", ")
	})
}

func PostTree(root *Node, f func(node *Node)) {
	if root == nil {
		return
	}
	PrintPostTree(root.Left)
	PrintPostTree(root.Right)
	f(root)
}
