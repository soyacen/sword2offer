package main

import (
	"fmt"
	"github.com/yacen/sword2offer/list"
	"strings"
)

/**
题目描述
牛客最近来了一个新员工Fish，每天早晨总是会拿着一本英文杂志，写些句子在本子上。
同事Cat对Fish写的内容颇感兴趣，有一天他向Fish借来翻看，但却读不懂它的意思。
例如，“student. a am I”。后来才意识到，这家伙原来把句子单词的顺序翻转了，
正确的句子应该是“I am a student.”。Cat对一一的翻转这些单词顺序可不在行，
你能帮助他么？
*/
func main() {
	fmt.Println(ReverseSentence("student. a am I"))
	fmt.Println(ReverseSentence2("student. a am I"))
	fmt.Println(ReverseSentence3("student. a am I"))
}

// split
func ReverseSentence(str string) (result string) {
	words := strings.Split(str, " ")
	for _, s := range words {
		result = s + " " + result
	}
	return
}

// 递归
func ReverseSentence2(str string) (result string) {
	i := strings.Index(str, " ")
	if i >= 0 {
		return ReverseSentence2(str[i+1:]) + " " + str[:i]
	}
	return str
}

// 栈
func ReverseSentence3(str string) (result string) {
	s := list.NewStack()
	i := -1
	for {
		i = strings.Index(str, " ")
		if i >= 0 {
			s.Push(str[:i])
			str = str[i+1:]
		} else {
			s.Push(str)
			break
		}
	}

	for s.Top() != nil {
		result += s.Pop().(string) + " "
	}

	return
}
