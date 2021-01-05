package main

import (
	"fmt"
	"time"
)

func canPartition(nums []int) bool {
	m := make(map[int]int)
	backtrace(nums, &m, 0)

	fmt.Println(m)
	for _, v := range m {
		if v > 1 {
			return true
		}
	}

	return false
}

func backtrace(nums []int, m *map[int]int, sum int) {
	if _, ok := (*m)[sum]; ok {
		(*m)[sum]++
		return
	}

	(*m)[sum]++

	for i := 1; i < len(nums); i++ {
		backtrace(nums[i:], m, sum+nums[i])
	}
}

func main() {
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
}
