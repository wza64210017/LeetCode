package main

import "fmt"

func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := (left + right) / 2
		if nums[middle] > middle {
			right--
		} else {
			left = middle + 1
		}
	}

	return left
}

func main() {
	//fmt.Println(missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9}))
	//fmt.Println(missingNumber([]int{0, 1, 3, 4, 5}))
	fmt.Println(missingNumber([]int{0,1}))
}
