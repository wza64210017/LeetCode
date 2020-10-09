package main

import (
	"fmt"
	"sort"
)

func isStraight(nums []int) bool {
	sort.Ints(nums)

	pos := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			continue
		}

		if nums[i] == nums[i+1] {
			return false
		}

		if pos == -1 {
			pos = i
		}
	}

	return nums[4]-nums[pos] < 5
}

func main() {
	fmt.Println(isStraight([]int{0, 0, 2, 2, 5}))
	fmt.Println(isStraight([]int{4, 7, 5, 9, 2}))
	fmt.Println(isStraight([]int{9, 2, 6, 7, 10}))
	fmt.Println(isStraight([]int{1, 2, 3, 4, 0}))
}
