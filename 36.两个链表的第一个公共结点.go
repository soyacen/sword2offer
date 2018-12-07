package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/util"
)

/**
题目描述
输入两个链表，找出它们的第一个公共结点。
*/
func main() {
	common := &ListNode{
		11, &ListNode{
			21, &ListNode{
				31, &ListNode{
					41, &ListNode{
						51, nil}}}}}
	head1 := &ListNode{
		1, &ListNode{
			2, &ListNode{
				3, &ListNode{
					4, &ListNode{
						5, common}}}}}
	head2 := &ListNode{
		10, &ListNode{
			20, &ListNode{
				30, &ListNode{
					40, common}}}}

	fmt.Println(FindFirstCommonNode(head1, head2) == common)
	fmt.Println(FindFirstCommonNode2(head1, head2) == common)
	fmt.Println(FindFirstCommonNode3(head1, head2) == common)

}

// 循环两个列表，比较两个节点是否相等， O(nm)
func FindFirstCommonNode(pHead1, pHead2 *ListNode) (node *ListNode) {
	for node := pHead1; node != nil; node = node.Next {
		for n := pHead2; n != nil; n = n.Next {
			if node == n {
				return node
			}
		}
	}
	return
}

// 两个链表有公共节点，只能是Y形状，应为至有1个Next指针，所以只能是前面不同，后面相同，
// 用两个栈结构放Node，最后顺序弹出Node，当两个不一样的，就返回这个Node的Next O(m+n+x)
func FindFirstCommonNode2(pHead1, pHead2 *ListNode) (node *ListNode) {
	stack1 := NewStack()
	for node := pHead1; node != nil; node = node.Next {
		stack1.Push(node)
	}
	stack2 := NewStack()
	for node := pHead2; node != nil; node = node.Next {
		stack2.Push(node)
	}
	sPop1 := stack1.Pop()
	sPop2 := stack2.Pop()

	for sPop1 == sPop2 {
		sPop1 = stack1.Pop()
		sPop2 = stack2.Pop()
	}

	return sPop1.(*ListNode).Next
}

// 县算出长度，长的那个先走遍历，当遍历到长度一样时候在一块遍历，直到遍历node相等。  O(m+n+x)
func FindFirstCommonNode3(pHead1, pHead2 *ListNode) (node *ListNode) {
	length1 := 0
	for node := pHead1; node != nil; node = node.Next {
		length1++
	}
	length2 := 0
	for node := pHead2; node != nil; node = node.Next {
		length2++
	}
	diff := length1 - length2
	startNode1 := pHead1
	startNode2 := pHead2
	if diff > 0 {
		for diff != 0 {
			startNode1 = startNode1.Next
			diff--
		}
	} else {
		for diff != 0 {
			startNode2 = startNode2.Next
			diff++
		}
	}
	node1 := startNode1
	node2 := startNode2
	for node1 != node2 {
		node1 = node1.Next
		node2 = node2.Next
	}
	return node1
}
