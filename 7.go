package main

import "fmt"

func reverse(x int) int {

	flag := false
	if x < 0 {
		flag = true
		x = -x
	}

	var num int64 = 0
	var num32 int64 = 1 << 31 - 1
	fmt.Println(num32)

	for x != 0 {
		num = num * 10 + int64(x % 10)

		if num > num32 {
			return 0
		}

		x /= 10
	}

	if flag {
		return -int(num)
	}

	return int(num)
}

func main() {
	fmt.Println(int(^uint32(0) >> 1))
	fmt.Println(reverse(120))
}
