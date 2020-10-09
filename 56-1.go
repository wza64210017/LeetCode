package main

import "fmt"

func singleNumbers(nums []int) []int {
	var a int
	for _, num := range nums {
		a ^= num
	}

	mask := a & (-a)
	ret := make([]int, 2)
	for _, num := range nums {
		if num&mask == 0 {
			ret[0] ^= num
		} else {
			ret[1] ^= num
		}
	}

	return ret
}

// 0000 0010    2
// 0000 1010    10
// 异或结果
// 0000 1000    8

// 求和任意数 与 都为0的数
// 0000 0100    4

func main() {
	arr := []int{1, 2, 10, 4, 1, 4, 3, 3}
	fmt.Println(singleNumbers(arr))
}
