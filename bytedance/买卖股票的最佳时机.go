package main

import "fmt"

func maxProfit(prices []int) int {
	min := 1 << 32
	max := 0

	for _, price := range prices {
		min = compare(min, price, "min")
		max = compare(max, price-min, "max")
	}

	return max
}

func compare(i, j int, t string) int {
	if i < j && t == "min" {
		return i
	} else if i < j && t == "max" {
		return j
	} else if i > j && t == "min" {
		return j
	} else {
		return i
	}
}

func main() {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
