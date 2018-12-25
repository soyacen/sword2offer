package main

import "fmt"

/**
题目描述
请实现一个函数用来找出字符流中第一个只出现一次的字符。
例如，当从字符流中只读出前两个字符"go"时，第一个只出现一次的字符是"g"。
当从该字符流中读出前六个字符“google"时，第一个只出现一次的字符是"l"。
输出描述:
如果当前字符流没有存在出现一次的字符，返回#字符。
*/
func main() {
	str := "google"
	solution := Solution1{stream: make([]rune, 0, 16)}
	for _, v := range str {
		solution.Insert(v)
		fmt.Println(string(v), string(solution.FirstAppearingOnce()))
	}

	fmt.Println("===================")
	solution2 := Solution2{stream: make([]rune, 0, 16)}
	for _, v := range str {
		solution2.Insert(v)
		fmt.Println(string(v), string(solution2.FirstAppearingOnce()))
	}

	fmt.Println("===================")
	solution3 := Solution3{stream: make([]rune, 0, 16), table: make(map[rune]int)}
	for _, v := range str {
		solution3.Insert(v)
		fmt.Println(string(v), string(solution3.FirstAppearingOnce()))
	}
}

// 时间 O（n2），空间O(n)
type Solution1 struct {
	stream []rune
}

//Insert one char from stringstream
func (s *Solution1) Insert(r rune) {
	for i, v := range s.stream {
		if v == r {
			s.stream = append(s.stream[:i], s.stream[i+1:]...)
			return
		}
	}
	s.stream = append(s.stream, r)
}

//return the first appearence once char in current stringstream
func (s *Solution1) FirstAppearingOnce() rune {
	if len(s.stream) == 0 {
		return '#'
	} else {
		return s.stream[0]
	}
}

// map存，时间O(n)，空间，O(n+256)
type Solution2 struct {
	stream []rune
	table  [256]int
}

//Insert one char from stringstream
func (s *Solution2) Insert(r rune) {
	s.table[r]++
	s.stream = append(s.stream, r)
}

//return the first appearence once char in current stringstream
func (s *Solution2) FirstAppearingOnce() rune {
	for _, v := range s.stream {
		if s.table[v] == 1 {
			return v
		}
	}
	return '#'
}

// map存，时间O(n)，空间，O(n)
type Solution3 struct {
	stream []rune
	table  map[rune]int
}

//Insert one char from stringstream
func (s *Solution3) Insert(r rune) {
	if _, ok := s.table[r]; ok {
		s.table[r]++
	} else {
		s.table[r] = 1
		s.stream = append(s.stream, r)
	}
}

//return the first appearence once char in current stringstream
func (s *Solution3) FirstAppearingOnce() rune {
	for _, v := range s.stream {
		if s.table[v] == 1 {
			return v
		}
	}
	return '#'
}
