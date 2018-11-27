package main

import (
	"fmt"
	"math"
)

// 给定一个double类型的浮点数base和int类型的整数exponent。求base的exponent次方。

func main() {
	for i := -10; i < 10; i++ {
		fmt.Println(Power(1.13, i) == math.Pow(1.13, float64(i)))
	}

}

func Power(base float64, exponent int) (result float64) {
	if base == 1 {
		return 1
	}
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	var exp int = exponent
	if exponent < 0 {
		exp = -exponent
	}

	tmp := base
	result = tmp
	exp--
	for exp > 0 {
		result = tmp * base
		tmp = result
		exp--
	}
	if exponent < 0 {
		result = 1 / result
	}
	return
}
