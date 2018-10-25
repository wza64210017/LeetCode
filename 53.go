package main

import "fmt"

func maxSubArray(nums []int) int {

	length := len(nums)
	if length == 0 {
		return 0
	}

	sum := make([]int, length)

	sum[0] = nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum[i-1] < 0 || nums[i]+sum[i-1] < 0 {
			sum[i] = nums[i]
		} else {
			sum[i] = nums[i] + sum[i-1]
		}


		if sum[i] > max {
			max = sum[i]
		}
	}

	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2,0, -1}))
}

//-2, 1, -3, 4, -1, 2, 1, -5, -4
