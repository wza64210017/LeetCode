package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	if length < 4 {
		return nil

	}
	sort.Ints(nums)
	//fmt.Println(nums)

	ret := make([][]int, 0, length*4)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			diff := target - nums[i] - nums[j]
			l := j + 1
			r := len(nums) - 1

			for l < r {
				if nums[l]+nums[r] < diff {
					l++
				} else if nums[l]+nums[r] > diff {
					r--
				} else {
					ret = append(ret, []int{nums[i], nums[j], nums[l], nums[r]})
					l, r = fourSumNext(nums, l, r)
				}
			}
		}
	}

	return ret
}

func fourSumNext(nums []int, l, r int) (int, int) {
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
	fmt.Println(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
	fmt.Println(fourSum([]int{0, 0, 0, 0}, 0))
	fmt.Println(fourSum([]int{-1, 0, 1, 2, -1, -4}, -1))
	fmt.Println(fourSum([]int{-1, -5, -5, -3, 2, 5, 0, 4}, -7))

}
