package main

import "fmt"

func exchange(nums []int) []int {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j]&1 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}

	return nums
}

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4, 5}))
}
