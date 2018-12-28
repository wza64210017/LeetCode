package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	length := len(nums)
	if length <= 2 {
		return nil
	}

	res := make([][]int, 0, length*3)
	sort.Ints(nums)

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, length-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if s == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l, r = next(nums, l, r)
			} else if s < 0 {
				l++
			} else if s > 0 {
				r--
			}
		}
	}

	return res
}

func next(nums []int, l int, r int) (int, int) {
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}

	return l, r
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
