package tree

type LinkNode struct {
	Data  interface{}
	Left  *LinkNode
	Right *LinkNode
	Next  *LinkNode
}

func (root *LinkNode) InTree(fn func(node *LinkNode)) {
	if root != nil {
		root.Left.InTree(fn)
		fn(root)
		root.Right.InTree(fn)
	}
}
