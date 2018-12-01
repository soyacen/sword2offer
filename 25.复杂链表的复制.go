package main

import (
	"fmt"
	"math/rand"
	"time"
)

//输入一个复杂链表（每个节点中有节点值，以及两个指针，一个指向下一个节点，另一个特殊指针指向任意一个节点），
// 返回结果为复制后复杂链表的head。（注意，输出结果中请不要返回参数中的节点引用，否则判题程序会直接返回空）

type RandomListNode struct {
	Data   int
	Next   *RandomListNode
	Random *RandomListNode
}

func main() {

	head := &RandomListNode{Data: 0}
	pre := head
	for i := 1; i < 10; i++ {
		curr := &RandomListNode{Data: i}
		pre.Next = curr
		pre = curr
	}
	rand.Seed(time.Now().UnixNano())
	for listNode := head; listNode != nil; listNode = listNode.Next {
		rand := rand.Intn(10)
		for node := head; node != nil; node = node.Next {
			if node.Data == rand {
				listNode.Random = node
				break
			}
		}
	}

	for node := head; node != nil; node = node.Next {
		fmt.Printf("%d,", node.Data)
		if node.Next != nil {
			fmt.Printf(" nexe = %d,", node.Next.Data)
		}
		if node.Random != nil {
			fmt.Printf(" random = %d", node.Random.Data)
		}
		fmt.Printf(" ====== %p ======= \n", node)
	}

	t1 := time.Now().UnixNano()
	cNode := Clone1(head)
	fmt.Println("time:", time.Now().UnixNano()-t1)
	for node := cNode; node != nil; node = node.Next {
		fmt.Printf("%d,", node.Data)
		if node.Next != nil {
			fmt.Printf(" nexe = %d,", node.Next.Data)
		}
		if node.Random != nil {
			fmt.Printf(" random = %d", node.Random.Data)
		}
		fmt.Printf(" ====== %p ======= \n", node)

	}

	t2 := time.Now().UnixNano()
	cNode2 := Clone2(head)
	fmt.Println("time:", time.Now().UnixNano()-t2)
	for node := cNode2; node != nil; node = node.Next {
		fmt.Printf("%d,", node.Data)
		if node.Next != nil {
			fmt.Printf(" nexe = %d,", node.Next.Data)
		}
		if node.Random != nil {
			fmt.Printf(" random = %d", node.Random.Data)
		}
		fmt.Printf(" ====== %p ======= \n", node)
	}

	t3 := time.Now().UnixNano()
	cNode3 := Clone3(head)
	fmt.Println("time:", time.Now().UnixNano()-t3)
	for node := cNode3; node != nil; node = node.Next {
		fmt.Printf("%d,", node.Data)
		if node.Next != nil {
			fmt.Printf(" nexe = %d,", node.Next.Data)
		}
		if node.Random != nil {
			fmt.Printf(" random = %d", node.Random.Data)
		}
		fmt.Printf(" ====== %p ======= \n", node)

	}

}

func Clone1(pHead *RandomListNode) (cloneHead *RandomListNode) {
	if pHead == nil {
		return
	}
	// 之考虑next，创建常规列表
	var newNode, pre *RandomListNode
	for node := pHead; node != nil; node = node.Next {
		newNode = &RandomListNode{Data: node.Data}
		if cloneHead == nil {
			cloneHead = newNode
			pre = newNode
		} else {
			pre.Next = newNode
			pre = newNode
		}
	}

	// 循环 处理 random
	node := pHead
	cNode := cloneHead
	for {
		if node == nil {
			break
		}

		for n := cloneHead; n != nil; n = n.Next {
			if n.Data == node.Random.Data {
				cNode.Random = n
				break
			}
		}
		node = node.Next
		cNode = cNode.Next
	}
	return
}

func Clone2(pHead *RandomListNode) (cloneHead *RandomListNode) {
	if pHead == nil {
		return
	}

	var newNode, pre *RandomListNode
	for node := pHead; node != nil; node = node.Next {

		// 检查 node数据是否以在  clonelist 里面
		if isIn, n := IsIn(cloneHead, node.Data); isIn {
			newNode = n
		} else {
			newNode = &RandomListNode{Data: node.Data}
		}

		if cloneHead == nil {
			cloneHead = newNode
		}

		// 检查 randomNode 数据是否以在  clonelist 里面
		if isIn, n := IsIn(cloneHead, node.Random.Data); isIn {
			newNode.Random = n
		} else {
			newNode.Random = &RandomListNode{Data: node.Random.Data}
		}

		if pre != nil {
			pre.Next = newNode
		}
		pre = newNode
	}

	return
}

func IsIn(pHead *RandomListNode, data int) (bool, *RandomListNode) {
	if pHead == nil {
		return false, nil
	}
	node := pHead
	for {
		if node == nil {
			break
		}
		if node.Data == data {
			return true, node
		}
		if node.Random != nil && node.Random.Data == data {
			return true, node.Random
		}
		node = node.Next
	}

	return false, nil
}

/*
在不使用辅助空间的情况下实现O(N)的时间效率。

第一步：根据原始链表的每个结点N创建对应的N'，然后将N‘通过pNext接到N的后面；

第二步：设置复制出来的结点的pSibling。假设原始链表上的N的pSibling指向结点S，那么其对应复制出来的N'是N->pNext指向的结点，同样S'也是结点S->pNext指向的结点。

第三步：把长链表拆分成两个链表，把奇数位置的结点用pNext连接起来的就是原始链表，把偶数位置的结点通过pNext连接起来的就是复制链表。

*/

func Clone3(pHead *RandomListNode) (cloneHead *RandomListNode) {
	if pHead == nil {
		return
	}
	for node := pHead; node != nil; node = node.Next {
		newNode := &RandomListNode{Data: node.Data}
		if cloneHead == nil {
			cloneHead = newNode
		}
		newNode.Next = node.Next
		node.Next = newNode
		node = newNode
	}

	for node := pHead; node != nil; node = node.Next {
		newNode := node.Next
		newNode.Random = node.Random.Next
		node = node.Next
	}

	for node := pHead; node != nil; node = node.Next {
		newNode := node.Next
		node.Next = newNode.Next
	}
	return
}
