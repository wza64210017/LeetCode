package main

import "fmt"

func majorityElement(nums []int) int {
	length := len(nums)
	m := make(map[int]int, length)
	pos := length / 2

	for _, v := range nums {
		m[v]++
		if m[v] > pos {
			return v
		}
	}

	return 0
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}
