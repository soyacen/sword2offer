package main

import "fmt"

// 请实现一个函数，将一个字符串中的每个空格替换成“%20”。
// 例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。

func main() {
	fmt.Println(replaceSpace("We Are Happy", "%20"))
}

func replaceSpace(origin, replaceValue string) (result string) {
	strRunes := []rune(origin)
	spaceRunes := []rune(" ")
	replaceRunes := []rune(replaceValue)
	resultRunes := make([]rune, 0)
	for _, r := range strRunes {
		if r != spaceRunes[0] {
			resultRunes = append(resultRunes, r)
		} else {
			resultRunes = append(resultRunes, replaceRunes...)
		}
	}
	result = string(resultRunes)
	return
}
