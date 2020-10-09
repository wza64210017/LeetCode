package main

import "fmt"

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))

	dp[0] = nums[0]
	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])

		if dp[i] > maxNum {
			maxNum = dp[i]
		}
	}

	return maxNum
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{-2, -1}))
}
