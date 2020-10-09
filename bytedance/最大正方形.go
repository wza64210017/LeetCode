package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	row, column := len(matrix), len(matrix[0])
	if row == 0 || column == 0 {
		return 0
	}

	dp := make([][]int, row+1)
	for i := 0; i <= row; i++ {
		dp[i] = make([]int, column+1)
	}

	ret := 0
	for i := 1; i <= row; i++ {
		for j := 1; j <= column; j++ {
			if matrix[i-1][j-1] == 1 {
				dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
			}

			ret = max(ret, dp[i][j]*dp[i][j])
		}
	}

	return ret
}

func min(nums ...int) int {
	min := 1 << 32
	for _, num := range nums {
		if min > num {
			min = num
		}
	}

	return min
}

func max(nums ...int) int {
	max := 0
	for _, num := range nums {
		if max < num {
			max = num
		}
	}

	return max
}

func main() {
	nums := [][]byte{
		//	{1, 0, 1, 0, 0},
		//	{1, 0, 1, 1, 1},
		//	{1, 1, 1, 1, 1},
		//	{1, 0, 0, 1, 0},
		{1},
	}

	fmt.Println(maximalSquare(nums))
}
