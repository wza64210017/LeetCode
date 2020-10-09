package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)
	result := make([][]int, 0)

	for now := 0; now < len(nums); now++ {
		if now > 0 && nums[now] == nums[now-1] {
			continue
		}

		left := now + 1
		right := len(nums) - 1
		for left < right {
			if nums[left]+nums[right] == -nums[now] {
				result = append(result, []int{nums[now], nums[left], nums[right]})
				left, right = next(nums, left, right)
			} else if nums[left]+nums[right] > -nums[now] {
				right--
			} else {
				left++
			}
		}
	}

	return result
}

func next(nums []int, left, right int) (int, int) {
	for left < right {
		if nums[left] == nums[left+1] {
			left++
		} else if nums[right] == nums[right-1] {
			right--
		} else {
			left++
			right--
			break
		}
	}

	return left, right
}

func main() {
	//fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}
