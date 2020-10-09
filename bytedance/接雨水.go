package main

import "fmt"

// 逐层计算总量
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	// 计算最大高度
	max := max(height)
	// 总量, 临时量, 左挡板
	sum, tmp, left := 0, 0, false

	for i := 1; i <= max; i++ {
		tmp, left = 0, false
		for _, val := range height {
			// 挡板大于水位
			if i > val {
				// 且有高于水位的左挡板
				if left {
					tmp++
				}

				continue
			}

			// 右挡板大于等于水位

			// 若无左挡板，将此处作为左挡板
			if left == false {
				left = true
				continue
			}

			// 有左挡板，且有右挡板，计入总水位
			sum += tmp
			tmp = 0
		}
	}

	return sum
}

// 计算出最大高度
func max(height []int) int {
	max := height[0]
	for _, val := range height {
		if val > max {
			max = val
		}
	}

	return max
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
	fmt.Println(trap([]int{1}))                                  // 0
	fmt.Println(trap([]int{5, 4, 1, 2}))                         // 1
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))                   // 9
	fmt.Println(trap([]int{7, 9, 7, 3}))                         // 0
	fmt.Println(trap([]int{4, 3, 3, 9, 3, 0, 9, 2, 8, 3}))       // 23
	fmt.Println(trap([]int{4, 9, 4, 5, 3, 2}))                   // 1
	fmt.Println(trap([]int{9, 6, 8, 8, 5, 6, 3}))                // 3
}

//			 *
// *         *    4
// *     *   *    3
// * *   * * *    1
// * *   * * *    1
