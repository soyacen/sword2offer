package main

import "fmt"

/*

在一个字符串(0<=字符串长度<=10000，全部由字母组成)中找到第一个只出现一次的字符,
并返回它的位置, 如果没有则返回 -1（需要区分大小写）.
*/
func main() {
	fmt.Println(FirstNotRepeatingChar("qq"))
}

type Char34 struct {
	Count      int
	FirstIndex int
}

func FirstNotRepeatingChar(str string) int {
	length := len(str)
	if length >= 0 && length <= 1000 {
		chMap := make(map[int32]*Char34)
		dupMap := make(map[int32]*Char34)
		for i, v := range str {
			if char, ok := chMap[v]; ok {
				char.Count++
				delete(chMap, v)
				dupMap[v] = char
			} else {
				if char, ok := dupMap[v]; ok {
					char.Count++
				}
				chMap[v] = &Char34{Count: 1, FirstIndex: i}
			}
		}
		if len(chMap) > 0 {
			minIndex := 0
			first := 0
			for _, c := range chMap {
				if first == 0 {
					first++
					minIndex = c.FirstIndex
				}
				if minIndex > c.FirstIndex {
					minIndex = c.FirstIndex
				}
			}
			return minIndex
		}

	}
	return -1
}
