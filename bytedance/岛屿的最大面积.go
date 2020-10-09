package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	ret := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				tmp := dfs(grid, i, j)
				ret = max(ret, tmp)
			}
		}
	}

	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func dfs(grid [][]int, row, col int) int {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return 0
	}

	if grid[row][col] != 1 {
		return 0
	}

	// 标记
	grid[row][col] = 2
	return 1 +
		dfs(grid, row, col-1) +
		dfs(grid, row, col+1) +
		dfs(grid, row-1, col) +
		dfs(grid, row+1, col)
}

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	fmt.Println(maxAreaOfIsland(grid))
}
