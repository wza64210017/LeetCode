package main

import (
	"fmt"
)

// 滑动窗口
func findContinuousSequence(target int) [][]int {
	if target < 3 {
		return nil
	}

	i, j, sum := 1, 1, 0
	res := make([][]int, 0)

	for i < target/2+1 {
		if sum < target {
			sum += j
			j++
		} else if sum > target {
			sum -= i
			i++
		} else {
			tmp := make([]int, 0, j-i)
			for t := i; t < j; t++ {
				tmp = append(tmp, t)
			}

			res = append(res, tmp)
			sum -= i
			i++
		}
	}

	return res
}

func main() {
	fmt.Println(findContinuousSequence(15))
}
