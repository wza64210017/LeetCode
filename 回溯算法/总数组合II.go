package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	ret := make([][]int, 0)
	backtrace(candidates, target, []int{}, &ret)

	return ret
}

func backtrace(candidates []int, target int, sub []int, ret *[][]int) {
	if target < 0 {
		return
	}

	if target == 0 {
		tmp := make([]int, len(sub))
		copy(tmp, sub)
		*ret = append(*ret, tmp)
		return
	}

	for i := 0; i < len(candidates); i++ {
		if i > 0 && candidates[i-1] == candidates[i] {
			continue
		}

		backtrace(candidates[i+1:], target-candidates[i], append(sub, candidates[i]), ret)
	}
}

func main() {
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
