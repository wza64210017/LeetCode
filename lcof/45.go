package main

import (
	"fmt"
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return toString(nums[i], nums[j]) < toString(nums[j], nums[i])
	})

	return toString(nums...)
}

func toString(nums ...int) string {
	s := ""
	for _, num := range nums {
		s += strconv.Itoa(num)
	}

	return s
}

func main() {
	//fmt.Println(minNumber([]int{3, 30, 34, 5, 9}))
	fmt.Println(minNumber([]int{10, 2}))
}
