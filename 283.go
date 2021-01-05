package main

func moveZeros(nums []int) {
	if len(nums) == 0 {
		return
	}

	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}

		nums[i], nums[p] = nums[p], nums[i]
		p++
	}
}

func main() {
	moveZeros([]int{-1, 1, 0, 3, 0, 12})
}
