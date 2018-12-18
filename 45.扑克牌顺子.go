package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/**
题目描述
LL今天心情特别好,因为他去买了一副扑克牌,发现里面居然有2个大王,2个小王(一副牌原本是54张^_^)...他随机从中抽出了5张牌,想测测自己的手气,
看看能不能抽到顺子,如果抽到的话,他决定去买体育彩票,嘿嘿！！“红心A,黑桃3,小王,大王,方片5”,“Oh My God!”不是顺子.....LL不高兴了,
他想了想,决定大\小 王可以看成任何数字,并且A看作1,J为11,Q为12,K为13。上面的5张牌就可以变成“1,2,3,4,5”(大小王分别看作2和4),“So Lucky!”。
LL决定去买体育彩票啦。 现在,要求你使用这幅牌模拟上面的过程,然后告诉我们LL的运气如何， 如果牌能组成顺子就输出true，否则就输出false。
为了方便起见,你可以认为大小王是0。
*/
func main() {
	pock := make([]int, 56)
	for i := 0; i < 4; i++ {
		for j := 1; j <= 13; j++ {
			pock[i*13+j-1] = j
		}
	}
	fmt.Println(pock)
	rand.Seed(time.Now().UnixNano())
	cards := make(map[int]int)
	for {
		index := rand.Intn(56)
		if _, ok := cards[index]; ok {
			continue
		} else {
			cards[index] = pock[index]
		}
		if len(cards) == 5 {
			break
		}
	}
	fmt.Println(cards)

	numbers := make([]int, 0, 5)
	for _, n := range cards {
		numbers = append(numbers, n)
	}
	fmt.Println(numbers)

	fmt.Println(IsContinuous(numbers))
	fmt.Println(IsContinuous([]int{1, 2, 3, 4, 5}))
	fmt.Println(IsContinuous([]int{1, 0, 3, 4, 5}))
	fmt.Println(IsContinuous([]int{1, 0, 0, 4, 5}))
	fmt.Println(IsContinuous([]int{1, 0, 0, 0, 0}))

	fmt.Println(IsContinuous([]int{9, 5, 8, 6, 7}))
	fmt.Println(IsContinuous([]int{12, 5, 8, 6, 7}))
	fmt.Println(IsContinuous([]int{12, 5, 8, 13, 7}))
	fmt.Println(IsContinuous([]int{9, 0, 8, 0, 7}))
}

func IsContinuous(numbers []int) bool {
	if len(numbers) < 5 {
		return false
	}
	sort.Ints(numbers)
	kings := 0
	pre := -1
	for _, v := range numbers {
		if v == 0 {
			kings++
			continue
		} else {
			if pre == -1 {
				pre = v
				continue
			} else {
				if pre == v {
					return false
				} else {
					if v-pre == 1 {
						pre = v
					} else {
						if v-pre > kings+1 {
							return false
						} else {
							kings = kings - (v - pre)
							kings = 0
							pre = v
						}
					}
				}
			}
		}
	}

	return true
}
