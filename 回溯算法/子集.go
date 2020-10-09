package main

import "fmt"

func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	backtrace(nums, []int{}, &ret)

	return ret
}

func backtrace(nums []int, sub []int, ret *[][]int) {
	tmp := make([]int, len(sub))
	copy(tmp, sub)
	*ret = append(*ret, tmp)

	for i := 0; i < len(nums); i++ {
		backtrace(nums[i+1:], append(sub, nums[i]), ret)
	}
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
