package main

import "fmt"

func maxSubArray(nums []int) int {
	ret := make([]int, len(nums))

	ret[0] = nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if ret[i-1] < 0 {
			ret[i] = nums[i]
		} else {
			ret[i] = ret[i-1] + nums[i]
		}

		if max < ret[i] {
			max = ret[i]
		}
	}

	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	//fmt.Println(maxSubArray([]int{1}))
}
