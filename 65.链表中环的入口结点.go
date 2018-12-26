package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/util"
)

/**
题目描述
给一个链表，若其中包含环，请找出该链表的环的入口结点，否则，输出null。
*/

/**'

1, 2, 3, 4, 5, 6,
		  11,	7,
		  10 9,	8,

*/
func main() {

	n4 := &ListNode{Data: 4}

	n1 := &ListNode{Data: 1, Next: &ListNode{Data: 2, Next: &ListNode{Data: 3, Next: n4}}}

	n4.Next = &ListNode{Data: 5, Next: &ListNode{Data: 6, Next: &ListNode{Data: 7, Next: &ListNode{Data: 8, Next: &ListNode{Data: 9, Next: &ListNode{Data: 10, Next: &ListNode{Data: 11, Next: n4}}}}}}}

	fmt.Println(EntryNodeOfLoop1(n1))
	fmt.Println(EntryNodeOfLoop2(n1))
	fmt.Println(EntryNodeOfLoop3(n1))
	fmt.Println(EntryNodeOfLoop4(n1))
}

// 遍历，同时放入map中，出现两次的就是环的入口 O(n)
func EntryNodeOfLoop1(pHead *ListNode) *ListNode {
	table := make(map[*ListNode]int)
	for n := pHead; n != nil; n = n.Next {
		if _, ok := table[n]; ok {
			return n
		} else {
			table[n] = 1
		}
	}
	return nil
}

/**
如果链表存在环，那么计算出环的长度n，
然后准备两个指针pSlow，pFast，pFast先走n步，
然后pSlow和pFase一块走，当两者相遇时，即为环的入口处；
s f  n
1 1  0
1 2  1
2 4  2
3 6  3
4 8  4
5 10 5
6 4  6
7 6  7
8 8  8

1 1  8
1 2  7
1 3  6
1 4  5
1 5  4
1 6  3
1 7  2
1 8  1
1 9  0
2 10 -1
3 11 -2
4 4  -3
*/
func EntryNodeOfLoop2(pHead *ListNode) *ListNode {
	pSlow, pFast := pHead, pHead
	n := 0
	for {
		n++
		pSlow = pSlow.Next
		if pSlow == nil {
			return nil
		}
		pFast = pFast.Next
		if pFast == nil {
			return nil
		}
		pFast = pFast.Next
		if pFast == nil {
			return nil
		}
		if pFast == pSlow {
			break
		}
	}
	pSlow, pFast = pHead, pHead

	for {
		n--
		pFast = pFast.Next
		if n < 0 {
			pSlow = pSlow.Next
		}
		if pSlow == pFast {
			return pSlow
		}
	}
}

/**
如果链表存在环，我们无需计算环的长度n，只需在相遇时，让一个指针在相遇点出发，
另一个指针在链表首部出发，然后两个指针一次走一步，当它们相遇时，就是环的入口处。
（这里就不说明为什么这样做是正确的，大家可以在纸上推导一下公式）
s f  n
1 1  0
1 2  1
2 4  2
3 6  3
4 8  4
5 10 5
6 4  6
7 6  7
8 8  8

1 9
2 10
3 11
4 4
*/

func EntryNodeOfLoop3(pHead *ListNode) *ListNode {
	pSlow, pFast := pHead, pHead
	n := 0
	for {
		n++
		pSlow = pSlow.Next
		if pSlow == nil {
			return nil
		}
		pFast = pFast.Next
		if pFast == nil {
			return nil
		}
		pFast = pFast.Next
		if pFast == nil {
			return nil
		}
		if pFast == pSlow {
			pSlow = pHead
			for pSlow != pFast {
				pSlow, pFast = pSlow.Next, pFast.Next
			}
			return pSlow
		}
	}
}

/*
使用断链法，在当前结点访问完毕后，断掉指向当前结点的指针。因此，最后一个被访问的结点一定是入口结点。
*/
func EntryNodeOfLoop4(pHead *ListNode) *ListNode {
	n := pHead
	for n.Next != nil {
		tmp := n
		n = n.Next
		tmp.Next = nil
	}
	return n
}
