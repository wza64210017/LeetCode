package main

import "fmt"

func plusOne(digits []int) []int {

	l := len(digits)
	arr := make([]int, l+1)
	add := 1
	for i := len(digits) - 1; i >= 0; i-- {
		tmp := digits[i] + add
		if tmp >= 10 {
			add = 1
			tmp = tmp - 10
		} else {
			add = 0
		}

		arr[i+1] = tmp
	}

	if add == 1 {
		arr[0] = 1
	}

	if arr[0] == 0 {
		arr = arr[1:]
	}

	return arr
}

func main() {
	fmt.Println(plusOne([]int{9}))
}
