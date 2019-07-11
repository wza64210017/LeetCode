package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	posMap := make(map[int]int)
	for pos, num := range nums2 {
		posMap[num] = pos
	}

	for i, num := range nums1 {
		locate := posMap[num]

		for _, value := range nums2[locate:] {
			if value > num {
				nums1[i] = value
				goto P
			}
		}

		nums1[i] = -1
	P:
	}

	return nums1
}

func main() {
	num1 := []int{4, 1, 2}
	num2 := []int{1, 3, 4, 2}

	fmt.Println(nextGreaterElement(num1, num2))
}
