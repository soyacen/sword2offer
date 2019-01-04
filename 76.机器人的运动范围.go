package main

import "fmt"

/**
题目描述
地上有一个m行和n列的方格。

一个机器人从坐标0,0的格子开始移动，每一次只能向左，右，上，下四个方向移动一格，
但是不能进入行坐标和列坐标的数位之和大于k的格子。

例如，当k为18时，机器人能够进入方格（35,37），因为3+5+3+7 = 18。

但是，它不能进入方格（35,38），因为3+5+3+8 = 19。

请问该机器人能够达到多少个格子？
*/
func main() {

	fmt.Println(movingCount(5, 10, 10), "[21]")
	fmt.Println(movingCount(15, 20, 20), "[359]")
	fmt.Println(movingCount(10, 1, 100), "[29]")
	fmt.Println(movingCount(10, 1, 10), "[10]")
	fmt.Println(movingCount(15, 100, 1), "[79]")
	fmt.Println(movingCount(15, 10, 1), "[10]")
	fmt.Println(movingCount(5, 10, 10), "[21]")
	fmt.Println(movingCount(12, 1, 1), "[1]")
	fmt.Println(movingCount(-10, 10, 10), "[0]")
}

/**
 广度优先算法遍历
0

0 0
0

0 0 0
0 0
0
*/
func movingCount(threshold, rows, cols int) int {
	bound := []int{rows, cols}
	// 保存所有步
	steps := [][]int{}

	// 存放当前探索出所有的步
	queue := [][]int{{0, 0}}

	// 上，左，下，右
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for len(queue) > 0 {
		// 走一步
		current := queue[0]
		queue = queue[1:]
		if outOfBounds(current, bound) {
			continue
		}
		if biggerThanThreshold(current, threshold) {
			continue
		}
		if alreadyPassed(current, steps) {
			continue
		}
		steps = append(steps, current)
		// 探索四个方向
		for _, dir := range dirs {
			newStep := []int{current[0] + dir[0], current[1] + dir[1]}
			queue = append(queue, newStep)
		}
	}
	return len(steps)
}

// 出界了
func outOfBounds(step []int, bound []int) bool {
	if step[0] < 0 || step[1] < 0 || step[0] >= bound[0] || step[1] >= bound[1] {
		return true
	}
	return false
}

// 大于 threshold 值
func biggerThanThreshold(step []int, threshold int) bool {
	i, j := step[0], step[1]
	k := 0
	for i > 0 {
		k += i % 10
		i /= 10
	}
	for j > 0 {
		k += j % 10
		j /= 10
	}
	if k > threshold {
		return true
	}

	return false
}

// 是否是以走过的路
func alreadyPassed(step []int, steps [][]int) bool {
	for _, item := range steps {
		if step[0] == item[0] && step[1] == item[1] {
			return true
		}
	}
	return false
}
