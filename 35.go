package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for k, num := range nums {
		if num == target || target < num{
			return k
		}
	}

	return len(nums)
}

func main() {
	nums := []int{1,3,5,6}
	target := 5

	fmt.Println(searchInsert(nums, target))
}
