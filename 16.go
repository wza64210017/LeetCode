package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	length := len(nums)
	diff := abs(nums[0] + nums[1] + nums[2] - target)
	res := nums[0] + nums[1] + nums[2]
	for i := range nums {
		l, r := i+1, length-1

		for l < r {
			tmp := nums[i] + nums[l] + nums[r]
			if tmp < target {
				l++
			} else if tmp > target {
				r--
			} else {
				return tmp
			}

			if abs(tmp-target) < diff {
				diff = abs(tmp - target)
				res = tmp
			}
		}
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	//fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{0, 1, 2}, 0))
}
