package main

import "fmt"

func maxValue(grid [][]int) int {
	width, length := len(grid), len(grid[0])

	m := grid[0][0]
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if i == 0 && j > 0 {
				grid[i][j] += grid[i][j-1]
			}

			if j == 0 && i > 0 {
				grid[i][j] += grid[i-1][j]
			}

			if i != 0 && j != 0 {
				grid[i][j] += max(grid[i][j-1], grid[i-1][j])
			}

			if m < grid[i][j] {
				m = grid[i][j]
			}
		}
	}

	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(maxValue([][]int{
		{1, 3 ,1},
		{1, 5, 1},
		{4, 2, 1},
	}))

	//fmt.Println(maxValue([][]int{
	//	{1},
	//	{2},
	//}))
}
