package main

import "fmt"

func twoSum(nums []int, target int) []int {

	m := make(map[int]int, len(nums))
	for k, num := range nums {
		m[num] = k
	}

	for k, v := range nums {
		diff := target-v
		if m[diff] != 0 && m[diff] != k {
			return []int{k, m[diff]}
		}
	}

	return nil
}

//两数之和
func main() {
	fmt.Println(twoSum([]int{2, 3, 11, 15}, 4))
}
