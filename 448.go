package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	arr := make([]int, 0, len(nums))
	for k, v := range nums {
		if k + 1 != v {
			arr = append(arr, k+1)
		}
	}

	return arr
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
