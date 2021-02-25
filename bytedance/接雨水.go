package main

import "fmt"

// 逐层计算水位
func trap(height []int) int {
	if len(height) < 2 {
		return 0
	}

	sum := 0
	for i := 1; i <= maxHeight(height); i++ {
		left, tmp := false, 0
		for _, v := range height {
			// 水位大于挡板高度
			if i > v {
				if left {
					tmp++
				}

				continue
			}

			// 左挡板未设置
			if left == false {
				left = true
				continue
			}

			// 存在右挡板
			sum += tmp
			tmp = 0
		}
	}

	return sum
}

// 最高挡板位置
func maxHeight(height []int) int {
	maxH := height[0]
	for _, v := range height {
		if v > maxH {
			maxH = v
		}
	}

	return maxH
}

func main() {
	//fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
	//fmt.Println(trap([]int{1}))                                  // 0
	//fmt.Println(trap([]int{5, 4, 1, 2}))                         // 1
	//fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))                   // 9
	//fmt.Println(trap([]int{7, 9, 7, 3}))                         // 0
	fmt.Println(trap([]int{4, 3, 3, 9, 3, 0, 9, 2, 8, 3})) // 23
	//fmt.Println(trap([]int{4, 9, 4, 5, 3, 2}))                   // 1
	//fmt.Println(trap([]int{9, 6, 8, 8, 5, 6, 3}))                // 3
}
