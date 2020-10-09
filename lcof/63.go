package main

import "fmt"

func maxProfit(prices []int) int {
	min := 1 << 32
	profit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}

		if prices[i]-min > profit {
			profit = prices[i] - min
		}
	}

	return profit
}

func main() {
	//fmt.Println(maxProfit([]int{}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
