package main

import (
	"fmt"
	"sort"
)

func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	backtrace(nums, []int{}, &ret)

	return ret
}

func backtrace(nums, sub []int, ret *[][]int) {
	tmp := make([]int, len(sub))
	copy(tmp, sub)
	sort.Ints(tmp)
	*ret = append(*ret, tmp)

	for k, num := range nums {
		backtrace(nums[k+1:], append(sub, num), ret)
	}
}

func main() {
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
}

//  [9 0 3 7]

//  [0,3,5,9]

// [[] [9] [9 0] [9 0 3] [9 0 3 7] [9 0 3 5 7] [9 0 3 7] [9 0 5] [9 0 5 7] [9 0 7] [9 3] [9 3 5] [9 3 5 7] [9 3 7] [9 5] [9 5 7] [9 7] [0] [0 3] [0 3 5] [0 3 5 7] [0 3 7] [0 5] [0 5 7] [0 7] [3] [3 5] [3 5 7] [3 7] [5] [5 7] [7]]
// [[],[9],[0],[0,9],[3],[3,9],[0,3],[0,3,9],[5],[5,9],[0,5],[0,5,9],[3,5],[3,5,9],[0,3,5],[0,3,5,9],[7],[7,9],[0,7],[0,7,9],[3,7],[3,7,9],[0,3,7],[0,3,7,9],[5,7],[5,7,9],[0,5,7],[0,5,7,9],[3,5,7],[3,5,7,9],[0,3,5,7],[0,3,5,7,9]]
