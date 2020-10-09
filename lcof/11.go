package main

import "fmt"

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left <= right {
		middle := (left + right) / 2
		if numbers[middle] > numbers[right] {
			left = middle + 1
		} else {
			right--
		}
	}

	return numbers[left]
}

func main() {
	fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
	// 3 4 5 0 1 2
	// l   m     r
	//     l   m r
	//     l m r
	//     l r
}
