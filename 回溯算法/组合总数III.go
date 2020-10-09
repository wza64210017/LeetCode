package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	ret := make([][]int, 0)
	backtrace(k, n, 1, []int{}, &ret)

	return ret
}

func backtrace(k, n, start int, sub []int, ret *[][]int) {
	if len(sub) > k {
		return
	}

	if n == 0 {
		tmp := make([]int, k)
		copy(tmp, sub)
		*ret = append(*ret, tmp)
		return
	}

	for i := start; i <= 9; i++ {
		if i > n {
			return
		}
		backtrace(k, n-i, i+1, append(sub, i), ret)
	}
}

func main() {
	fmt.Println(combinationSum3(3, 9))
}
