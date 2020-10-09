package main

import "fmt"

// 输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。
// 如果有多对数字的和等于s，则输出任意一对即可。
//func twoSum(nums []int, target int) []int {
//	for i := 0; i < len(nums)-1; i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i]+nums[j] == target {
//				return []int{nums[i], nums[j]}
//			}
//		}
//	}
//
//	return nil
//}

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] == target {
			break
		} else if nums[left]+nums[right] > target {
			right--
		} else {
			left++
		}
	}

	return []int{nums[left], nums[right]}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
