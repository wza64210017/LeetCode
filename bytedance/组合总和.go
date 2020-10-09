package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	ret := make([][]int, 0)
	backtrace(0, target, candidates, []int{}, &ret)

	return ret
}

func backtrace(start, target int, nums []int, sub []int, ret *[][]int) {
	if target < 0 {
		return
	}

	if target == 0 {
		tmp := make([]int, len(sub))
		copy(tmp, sub)
		*ret = append(*ret, tmp)
		return
	}

	for i := start; i < len(nums); i++ {
		backtrace(i, target-nums[i], nums, append(sub, nums[i]), ret)
	}
}

func main() {
	fmt.Println(combinationSum([]int{21, 46, 35, 20, 44, 31, 29, 23, 45, 37, 33, 34, 39, 42, 24, 40, 41, 26, 22, 38, 36, 27, 25, 49, 48, 43}, 71))
}
