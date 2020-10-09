package main

import "fmt"

// 数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
// 超过数组一半长，这个数字必定相连
func majorityElement(nums []int) int {
	val, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			val = nums[i]
		}

		if val == nums[i] {
			count++
		} else {
			count--
		}
	}

	return val
}

func main() {
	nums := []int{6, 5, 5}
	fmt.Println(majorityElement(nums))
}
