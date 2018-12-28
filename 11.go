package main

import "fmt"

func maxArea(height []int) int {
	max := 0
	p1 := 0
	p2 := len(height) - 1

	for p1 < p2 {
		l := Min(height[p1], height[p2])
		max = Max(max, l*(p2-p1))
		if height[p1] < height[p2] {
			p1++
		} else {
			p2--
		}
	}

	return max
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func Max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func main() {
	//fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))
}

//1: 2 * 6 = 12  p1: 1, p2: 6
//2: 3 * 5 = 15  p1: 2, p2: 6
//3: 4 * 4 = 16  p1: 3, p2: 6
//4: 5 * 3 = 16  p1: 3, p2: 17
//5: 5 * 2 = 10  p1: 3, p2: 18

//1: 2 * 6 = 12  p1: 1, p2: 6
//2: 3 * 5 = 15  p1: 2, p2: 6
//3: 4 * 4 = 16  p1: 3, p2: 6
//4: 5 * 3 = 16  p1: 4, p2: 5
//5: 17 * 1 = 17 p1: 5, p2: 5
