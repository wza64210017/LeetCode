package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}

	for i := 1; i < len(triangle); i++ {
		for j := 0; j < i; j++ {
			tmp := triangle[i-1][j]+triangle[i][j]
			if j >= 1 {
				tmp = min(tmp, triangle[i-1][j-1]+triangle[i][j])
			}

			triangle[i][j] = tmp
		}

		triangle[i][i] += triangle[i-1][i-1]
	}

	ret := triangle[len(triangle)-1][0]
	for _, v := range triangle[len(triangle)-1] {
		if v < ret {
			ret = v
		}
	}

	return ret
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func main() {
	nums := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}

	nums = [][]int{
		{-1},
		{2, 3},
		{1, -1, -3},
	}

	fmt.Println(minimumTotal(nums))
}
