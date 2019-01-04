package main

import (
	"fmt"
	"github.com/Workiva/go-datastructures/queue"
)

/**
题目描述
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，
那么中位数就是所有数值排序之后位于中间的数值。
如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。
我们使用Insert()方法读取数据流，使用GetMedian()方法获取当前读取数据的中位数。
*/
func main() {
	// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
	dataStream := []int{6, 1, 7, 2, 9, 3, 8, 4, 5, 0}

	s1 := &Solution72_1{}

	for _, v := range dataStream {
		s1.Insert(v)
	}

	fmt.Println(s1.GetMedian())

	s2 := &Solution72_2{}

	for _, v := range dataStream {
		s2.Insert(v)
	}

	fmt.Println(s2.GetMedian())
}

// 插入排序
type Solution72_1 struct {
	Data []int
}

func (s *Solution72_1) Insert(num int) {
	s.Data = append(s.Data, num)
	length := len(s.Data)
	for i := length - 1; i > 0; i-- {
		if s.Data[i] < s.Data[i-1] {
			s.Data[i], s.Data[i-1] = s.Data[i-1], s.Data[i]
		} else {
			break
		}
	}
}

func (s *Solution72_1) GetMedian() int {
	length := len(s.Data)
	if length%2 == 0 {
		return (s.Data[length/2] + s.Data[length/2-1]) / 2
	} else {
		return s.Data[length/2]
	}
}

// 大顶堆，小顶堆
type Solution72_2 struct {
	minHeap *queue.PriorityQueue
	maxHeap *queue.PriorityQueue
	count   int
}

type MinItem int

func (mi MinItem) Compare(other queue.Item) int {
	omi := other.(MinItem)
	if mi > omi {
		return 1
	} else if mi == omi {
		return 0
	}
	return -1
}

type MaxItem int

func (mi MaxItem) Compare(other queue.Item) int {
	omi := other.(MaxItem)
	if mi < omi {
		return 1
	} else if mi == omi {
		return 0
	}
	return -1
}

func (s *Solution72_2) Insert(num int) {
	if s.minHeap == nil {
		s.minHeap = queue.NewPriorityQueue(1, true)
	}
	if s.maxHeap == nil {
		s.maxHeap = queue.NewPriorityQueue(1, true)
	}
	if s.count%2 == 0 {
		//当数据总数为偶数时，新加入的元素，应当进入小根堆
		//（注意不是直接进入小根堆，而是经大根堆筛选后取大根堆中最大元素进入小根堆）
		//1.新加入的元素先入到大根堆，由大根堆筛选出堆中最大的元素
		err := s.maxHeap.Put(MaxItem(num))
		if err != nil {
			panic(err)
		}
		items, err := s.maxHeap.Get(1)
		if err != nil {
			panic(err)
		}
		n := int(items[0].(MaxItem))
		err = s.minHeap.Put(MinItem(n))
		if err != nil {
			panic(err)
		}
	} else {
		//当数据总数为奇数时，新加入的元素，应当进入大根堆
		//（注意不是直接进入大根堆，而是经小根堆筛选后取小根堆中最大元素进入大根堆）
		//1.新加入的元素先入到小根堆，由小根堆筛选出堆中最小的元素
		err := s.minHeap.Put(MinItem(num))
		if err != nil {
			panic(err)
		}
		items, err := s.minHeap.Get(1)
		if err != nil {
			panic(err)
		}
		var n int = int(items[0].(MinItem))
		err = s.maxHeap.Put(MaxItem(n))
		if err != nil {
			panic(err)
		}
	}
	s.count++
}

func (s *Solution72_2) GetMedian() int {
	if s.count%2 == 1 {
		return int(s.minHeap.Peek().(MinItem))
	} else {
		return (int(s.minHeap.Peek().(MinItem)) + int(s.maxHeap.Peek().(MaxItem))) / 2
	}
}
