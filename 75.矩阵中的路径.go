package main

import "fmt"

/**
题目描述
请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。

路径可以从矩阵中的任意一个格子开始，每一步可以在矩阵中向左

，向右，向上，向下移动一个格子。如果一条路径经过了矩阵中的某一个格子，

则之后不能再次进入这个格子。

例如
a b c e
s f c s
a d e e 这样的3 X 4 矩阵中

包含一条字符串"bcced"的路径，但是矩阵中不包含"abcb"路径，
因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，
路径不能再次进入该格子。
*/
func main() {
	fmt.Println(hasPath("abcesfcsadee", 3, 4, "bcced"))
	fmt.Println(hasPath("abcesfcsadee", 3, 4, "abcb"))
	fmt.Println(hasPath("abcesfcsadee", 3, 4, "asfdecs"))
	fmt.Println(hasPath("abcesfcsadee", 3, 4, "ade"))
	fmt.Println(hasPath("abcesfcsadee", 3, 4, "asa"))
	fmt.Println(hasPath("abcesfcsadee", 3, 4, "asadfbcced"))
}

func hasPath(matrix string, rows, cols int, str string) bool {
	strLen := len(str)
	if strLen == 0 {
		return true
	}
	matrixLen := len(matrix)
	if matrixLen == 0 {
		return false
	}
	if strLen > matrixLen {
		return false
	}
	if matrixLen != rows*cols {
		return false
	}

	runeMatrix := []rune(matrix)
	runeStr := []rune(str)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if p := check(i, j, rows, cols, runeMatrix, runeStr); p != nil {
				return true
			}
		}
	}

	return false
}

func check(i int, j int, rows, cols int, matrix []rune, str []rune) []int {
	if len(str) == 0 {
		return []int{}
	}
	if matrix[i*cols+j] != str[0] {
		return nil
	} else {
		if i-1 >= 0 {
			path := check(i-1, j, rows, cols, matrix, str[1:])
			if path != nil {
				if !existsAtPath(path, i, j) {
					path = append(path, i)
					path = append(path, j)
					return path
				}
			}
		}
		if i+1 < rows {
			path := check(i+1, j, rows, cols, matrix, str[1:])
			if path != nil {
				if !existsAtPath(path, i, j) {
					path = append(path, i)
					path = append(path, j)
					return path
				}
			}
		}
		if j-1 >= 0 {
			path := check(i, j-1, rows, cols, matrix, str[1:])
			if path != nil {
				if !existsAtPath(path, i, j) {
					path = append(path, i)
					path = append(path, j)
					return path
				}
			}
		}
		if j+1 < cols {
			path := check(i, j+1, rows, cols, matrix, str[1:])
			if path != nil {
				if !existsAtPath(path, i, j) {
					path = append(path, i)
					path = append(path, j)
					return path
				}
			}
		}
		return nil
	}
}

func existsAtPath(path []int, i int, j int) bool {
	for p := 0; p < len(path); p += 2 {
		if path[p] == i && path[p+1] == j {
			return true
		}
	}
	return false
}
