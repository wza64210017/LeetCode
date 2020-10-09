package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tmp, max := 1, 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			tmp = 1
			continue
		}

		tmp++
		if tmp > max {
			max = tmp
		}
	}

	return max
}

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 2, 0, 5, 7, 8, 9, 10}))
}
