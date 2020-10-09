package main

import "fmt"

func search(nums []int, target int) int {
	return find(nums, target) - find(nums, target-1)
}

func find(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] <= target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return left
}

func main() {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
	//fmt.Println(search([]int{1}, 1))
}
