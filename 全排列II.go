package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	ret := make([][]int, 0)
	use := make([]bool, len(nums))
	sub := make([]int, 0)
	backtrace(nums, sub, use, &ret)

	return ret
}

func backtrace(nums []int, sub []int, used []bool, ret *[][]int) {
	if len(sub) == len(nums) {
		tmp := make([]int, len(sub))
		copy(tmp, sub)
		*ret = append(*ret, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		used[i] = true
		backtrace(nums, append(sub, nums[i]), used, ret)
		used[i] = false
	}
}

func main() {
	fmt.Println(permuteUnique([]int{-1, 2, -1, 2, 1, -1, 2, 1}))
}
