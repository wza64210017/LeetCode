package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		if m[num] > 0 {
			return true
		}
		m[num]++
	}

	return false
}

func main() {
	nums := []int{1,2,3,4}
	fmt.Println(containsDuplicate(nums))
}
