package main

import "fmt"

func maxProfit(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}

	sum, pos := 0, 0
	for i := 1; i < length; i++ {
		if prices[i] < prices[i-1] {
			sum += prices[i-1] - prices[pos]
			pos = i
		}
	}

	if prices[length-1] > prices[pos] {
		sum += prices[length-1] - prices[pos]
	}

	return sum
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
