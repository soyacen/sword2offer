package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/util"
)

/**
题目描述
在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，
重复的结点不保留，返回链表头指针。 例如，链表1->2->3->3->4->4->5
处理后为 1->2->5
*/
func main() {
	list := &ListNode{Data: 1, Next: &ListNode{Data: 2, Next: &ListNode{Data: 3, Next: &ListNode{Data: 3, Next: &ListNode{Data: 4, Next: &ListNode{Data: 4, Next: &ListNode{Data: 5}}}}}}}

	n1 := deleteDuplication1(list)

	for n := n1; n != nil; n = n.Next {
		fmt.Println(n.Data)
	}
	fmt.Println("===============================")
	n1 = deleteDuplication1(&ListNode{Data: 1, Next: list})

	for n := n1; n != nil; n = n.Next {
		fmt.Println(n.Data)
	}

	fmt.Println("===============================")

	n1 = deleteDuplication2(list)

	for n := n1; n != nil; n = n.Next {
		fmt.Println(n.Data)
	}
	fmt.Println("===============================")
	n1 = deleteDuplication2(&ListNode{Data: 1, Next: list})

	for n := n1; n != nil; n = n.Next {
		fmt.Println(n.Data)
	}

	fmt.Println("===============================")

	n1 = deleteDuplication3(list)

	for n := n1; n != nil; n = n.Next {
		fmt.Println(n.Data)
	}
	fmt.Println("===============================")
	n1 = deleteDuplication3(&ListNode{Data: 1, Next: list})

	for n := n1; n != nil; n = n.Next {
		fmt.Println(n.Data)
	}
}

// O(n)
func deleteDuplication1(pHead *ListNode) *ListNode {
	// 前一个节点， 新列表节点，新列表尾节点， 尾节点的前一个节点
	var pre, newList, tail, preTail *ListNode
	for n := pHead; n != nil; n = n.Next {
		if pre == nil {
			// 第一个元素，初始化新列表
			pre = n
			newList = n
			tail = newList
		} else {
			if pre.Data != n.Data {
				// 当前一个节点的值与当前节点的值不同，

				// 如果当前雷表为空，初始化列表
				if newList == nil {
					newList = n
					tail = n
					preTail = nil
				} else {
					// 把尾节点赋值给前一个尾节点
					preTail = tail
					// 加上新的节点
					tail.Next = n
					tail = n
				}
				pre = n
			} else {
				// 前一个节点的值 与 当前节点的值 相同
				if tail.Data == n.Data {
					// 为节点 与 当前节点 相同
					if preTail != nil {
						// 前一个尾节点不为空， 就删掉为节点
						preTail.Next = nil
						tail = preTail
					} else {
						// 前一个尾节点为空， 新列表只有一个元素，开头就重复
						newList = nil
						tail = nil
					}
				}
				pre = n
			}
		}
	}
	return newList
}

// 改进版，好理解
func deleteDuplication2(pHead *ListNode) *ListNode {
	newList := &ListNode{}
	var pre, tail, preTail *ListNode
	for n := pHead; n != nil; n = n.Next {
		if pre == nil {
			// 第一个元素，初始化新列表
			newList.Next = n
			tail = n
			preTail = newList
		} else {
			if pre.Data != n.Data {
				// 当前一个节点的值与当前节点的值不同，

				// 如果当前雷表为空，初始化列表
				if newList.Next == nil {
					newList.Next = n
					tail = n
					preTail = newList
				} else {
					// 把尾节点赋值给前一个尾节点
					preTail = tail
					// 加上新的节点
					tail.Next = n
					tail = n
				}

			} else {
				// 前一个节点的值 与 当前节点的值 相同

				if tail.Data == n.Data {
					// 尾节点 与 当前节点 相同

					// 倒退尾节点
					preTail.Next = nil
					tail = preTail
				}
			}
		}
		pre = n
	}
	// 返回下一个
	return newList.Next
}

// 递归版
func deleteDuplication3(pHead *ListNode) *ListNode {
	if pHead == nil {
		return nil
	}
	n := pHead
	for ; n != nil; n = n.Next {
		if n.Next != nil {
			if n.Data != n.Next.Data {
				break
			}
		} else {
			break
		}
	}

	if n == pHead {
		n.Next = deleteDuplication3(n.Next)
		return n
	} else {
		return deleteDuplication3(n.Next)
	}
}
