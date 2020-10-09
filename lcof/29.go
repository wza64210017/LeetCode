package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	res := make([]int, 0, (r+1)*(b+1))
	for {
		// 向右移动
		for i := l; i <= r; i++ {
			res = append(res, matrix[t][i])
		}

		t++
		if t > b {
			break
		}

		// 向下移动
		for i := t; i <= b; i++ {
			res = append(res, matrix[i][r])
		}

		r--
		if l > r {
			break
		}

		// 向左移动
		for i := r; i > l-1; i-- {
			res = append(res, matrix[b][i])
		}

		b--
		if t > b {
			break
		}

		// 向上移动
		for i := b; i > t-1; i-- {
			res = append(res, matrix[i][l])
		}

		l++
		if l > r {
			break
		}
	}

	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}))
}
