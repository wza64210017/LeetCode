package main

import "fmt"

func removeElement(nums []int, val int) int {

	pos := 0
	for _, num := range nums {
		if num != val {
			nums[pos] = num
			pos++
		}
	}

	return pos
}

//移除元素
func main() {
	nums := []int{0,1,2,2,3,0,4,2}
	fmt.Println(removeElement(nums, 2))
	fmt.Println(nums)
}
