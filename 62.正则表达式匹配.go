package main

import (
	"fmt"
	"regexp"
)

/**
题目描述
请实现一个函数用来匹配包括'.'和'*'的正则表达式。模式中的字符'.'表示任意一个字符，
而'*'表示它前面的字符可以出现任意次（包含0次）。 在本题中，
匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，
但是与"aa.a"和"ab*a"均不匹配
*/
func main() {

	fmt.Println(regexp.MatchString("a*a", "aaaaa"))

	fmt.Println(match("aaa", "a.a"))
	fmt.Println(match("aaa", "ab*ac*a"))
	fmt.Println(match("aaa", "aa.a"))
	fmt.Println(match("aaa", "ab*a"))
}

func match(str, pattern string) bool {
	// 如果字符串和正则都为空串，则匹配成功
	if len(str) == 0 && len(pattern) == 0 {
		return true
	}
	// 如果字符号不为空，正则为空，匹配失败
	if len(str) > 0 && len(pattern) == 0 {
		return false
	}
	// 如果字符串为空，正则不为空，只有当正则为 字符+*，才匹配成功，其他都失败
	if len(str) == 0 && len(pattern) > 0 {
		// 如果长度奇数，匹配必失败
		if len(pattern)%2 == 1 {
			return false
		}
		for i := 0; i < len(pattern); i += 2 {
			// 正则有 ** 时，不合法
			if pattern[i] == '*' {
				return false
			}
			// 如果不是*， 匹配失败
			if pattern[i] != '*' {
				return false
			}
		}
		return true
	}
	// 如果字符串个正则都不为空
	if str[0] == pattern[0] || pattern[0] == '.' {
		// 如果正则第二个字符为 * ,就要匹配多个
		if len(pattern) > 1 && pattern[1] == '*' {
			// 如果出现类似与 p*p..p*..p 这中情况
			i := 2
			patCount := 0
			for {
				if pattern[i] == str[0] || pattern[i] == '.' {
					if i+1 < len(pattern) && pattern[i+1] == '*' {
						i += 2
					} else {
						patCount++
					}

				} else {
					break
				}
			}

			// 计算连续出现多少个相同的字符
			realCount := 1
			for ; str[realCount] == str[0]; realCount++ {
			}

			if patCount > realCount {
				// 如果正则出现的最少字符数大于字符串连续出现的字符，匹配失败
				return false
			} else {
				// 继续匹配剩下的
				return match(str[realCount:], pattern[i:])
			}

		} else {
			return match(str[1:], pattern[1:])
		}

	} else {
		// 如果正则的第二个字符不是 * ,匹配必失败，是 * 就继续匹配
		if len(pattern) > 1 {
			if pattern[1] != '*' {
				return false
			} else {
				return match(str, pattern[2:])
			}
		} else {
			return false
		}
	}
}
