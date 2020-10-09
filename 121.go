package main

import "fmt"

func maxProfit(prices []int) int {
	minNum := 1 << 32
	maxNum := 0

	for _, price := range prices {
		minNum = min(minNum, price)
		maxNum = max(maxNum, price - minNum)
	}

	return maxNum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
