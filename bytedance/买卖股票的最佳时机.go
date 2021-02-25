package main

import "fmt"

func maxProfit(prices []int) int {
	minNum := 1 << 32
	maxNUm := 0

	for _, price := range prices {
		minNum = min(minNum, price)
		maxNUm = max(maxNUm, price-minNum)
	}

	return maxNUm
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func main() {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
