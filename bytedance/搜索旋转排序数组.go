package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := (left + right) / 2

		switch {
		case target == nums[middle]:
			return middle
		case target == nums[left]:
			return left
		case target == nums[right]:
			return right
		case nums[middle] < nums[right]:
			if target > nums[middle] && target <= nums[right] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		case nums[middle] >= nums[left]:
			if target < nums[middle] && target >= nums[left] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{1, 3}, 3))
	fmt.Println(search([]int{1}, 1))
	fmt.Println(search([]int{1, 3}, 0))
	fmt.Println(search([]int{1}, 1))
}
