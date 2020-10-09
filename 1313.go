package main

import "fmt"

func decompressRLElist(nums []int) []int {
	ret := make([]int, 0, len(nums)*2)
	for i := 0; i < len(nums); i += 2 {
		for nums[i] > 0 {
			ret = append(ret, nums[i+1])
			nums[i]--
		}
	}

	return ret
}

func main() {
	fmt.Println(decompressRLElist([]int{1,1,2,3}))
}
