package main

import "fmt"

func combine(n int, k int) [][]int {
	nums := make([]int, 0)
	ret := make([][]int, 0)
	backTrace(1, n, k, nums, &ret)

	return ret
}

func backTrace(start, n, k int, nums []int, ret *[][]int) {
	if len(nums) == k {
		tmp := make([]int, k)
		copy(tmp, nums)
		*ret = append(*ret, tmp)
		return
	}

	for i := start; i <= n-(k-len(nums))+1; i++ {
		backTrace(i+1, n, k, append(nums, i), ret)
	}
}

func main() {
	fmt.Println(combine(13, 13))
}
