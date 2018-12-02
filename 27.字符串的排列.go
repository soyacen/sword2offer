package main

import "fmt"

/**
题目描述
输入一个字符串,按字典序打印出该字符串中字符的所有排列。例如输入字符串abc,则打印出由字符a,b,c所能排列出来的所有字符串abc,acb,bac,bca,cab和cba。
输入描述:
输入一个字符串,长度不超过9(可能有字符重复),字符只包括大小写字母。
*/
func main() {

	p1 := Permutation("abc")
	for i, s := range p1 {
		fmt.Println(i+1, s)
	}

	/*
		p2 := Permutation("qwertyuio")
		for i, s := range p2 {
			fmt.Println(i+1, s)
		}
	*/
	/*
		p3 := Permutation("aabbcc")
			for i, s := range p3 {
				fmt.Println(i+1, s)
			}
	*/

	/*	p4 := Permutation("csasdwqwrwqffwq")
		for i, s := range p4 {
			fmt.Println(i+1, s)
		}

		p5 := Permutation("csd234")
		for i, s := range p5 {
			fmt.Println(i+1, s)
		}*/
}

func Permutation(str string) (result []string) {
	if len(str) > 9 {
		return nil
	}
	for _, c := range str {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			continue
		} else {
			return nil
		}
	}
	return PermutationBytes(nil, []byte(str))
}

func PermutationBytes(pre, rest []byte) (result []string) {
	restLength := len(rest)
	if restLength == 0 {
		return []string{string(pre)}
	}
	var uniqRest []byte

	for i, c := range rest {
		uniq := true
		for _, uc := range uniqRest {
			if c == uc {
				uniq = false
				break
			}
		}
		if !uniq {
			continue
		}
		uniqRest = append(uniqRest, c)

		p := append(pre, c)
		r := make([]byte, restLength-1)
		rIndex := 0
		for rIndex < restLength {
			if rIndex < i {
				r[rIndex] = rest[rIndex]
			} else if rIndex > i {
				r[rIndex-1] = rest[rIndex]
			}
			rIndex++
		}
		result = append(result, PermutationBytes(p, r)...)
	}
	return
}
