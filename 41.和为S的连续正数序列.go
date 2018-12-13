package main

import "fmt"

/**
题目描述
小明很喜欢数学,有一天他在做数学作业时,要求计算出9~16的和,
他马上就写出了正确答案是100。但是他并不满足于此,
他在想究竟有多少种连续的正数序列的和为100(至少包括两个数)。
没多久,他就得到另一组连续正数和为100的序列:18,19,20,21,22。
现在把问题交给你,你能不能也很快的找出所有和为S的连续正数序列?
Good Luck!
输出描述:
输出所有和为S的连续正数序列。序列内按照从小至大的顺序，
序列间按照开始数字从小到大的顺序
*/
func main() {
	fmt.Println(FindContinuousSequence1(100))
	fmt.Println(FindContinuousSequence2(100))
	fmt.Println(FindContinuousSequence3(100))
}

func FindContinuousSequence1(sum int) (result [][]int) {
	tmp := make([][]int, 0)
	for n := 1; n <= sum/2+1; n++ {
		// 循环数组，判断首元素是否等于sum，等于的赋值为1代表成功，大于的就赋值为0代表失败，其他就继续相加
		for i, _ := range tmp {
			if tmp[i][0] > 0 {
				tmp[i] = append(tmp[i], n)
				tmp[i][0] += n
			}
			if tmp[i][0] == sum {
				tmp[i][0] = -1
			} else if tmp[i][0] > sum {
				tmp[i][0] = 0
			}
		}
		// 新加入一个数组，首元素为数组的和，第二个元素是第一个数字
		tmp = append(tmp, []int{n, n})
	}
	// 整理数组
	for i, _ := range tmp {
		if tmp[i][0] == -1 {
			result = append(result, tmp[i][1:])
		}
	}
	return
}

func FindContinuousSequence2(sum int) (result [][]int) {
	tmp := make([][]int, 0)
	for n := 1; n <= sum/2+1; n++ {
		// 循环数组，判断首元素是否等于sum，等于的赋值为1代表成功，大于的就赋值为0代表失败，其他就继续相加
		for i, _ := range tmp {
			if len(tmp[i]) > 0 {
				tmp[i] = append(tmp[i], n)
				tmpSum := (tmp[i][0] + tmp[i][len(tmp[i])-1]) * len(tmp[i]) / 2
				if tmpSum == sum {
					result = append(result, tmp[i])
					tmp[i] = nil
				} else if tmpSum > sum {
					tmp[i] = nil
				}
			}
		}
		// 新加入一个数组
		tmp = append(tmp, []int{n})
	}
	return
}

/**
滑动窗口法
链接：https://www.nowcoder.com/questionTerminal/c451a3fd84b64cb19485dad758a55ebe
来源：牛客网

用两个数字begin和end分别表示序列的最大值和最小值，
首先将begin初始化为1，end初始化为2.
如果从begin到end的和大于s，我们就从序列中去掉较小的值(即增大begin),
相反，只需要增大end。
如果和等于s，则记录begin到end的数组，begin+=1,end=begin+1,在开始
终止条件为：一直增加begin到(1+sum)/2并且end小于sum为止
*/
func FindContinuousSequence3(sum int) (result [][]int) {
	begin := 1
	end := 2
	for begin < (1+sum)/2 && end < sum {
		tmpSum := (begin + end) * (end - begin + 1) / 2
		if tmpSum == sum {
			tmpArr := make([]int, end-begin+1)
			for i, _ := range tmpArr {
				tmpArr[i] = begin + i
			}
			result = append(result, tmpArr)
			begin++
			end = begin + 1
		} else if tmpSum > sum {
			begin++
		} else {
			end++
		}
	}
	return
}
