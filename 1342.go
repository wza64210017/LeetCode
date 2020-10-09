package main

import "fmt"

func numberOfSteps(num int) int {
	count := 0
	for num > 0 {
		count++
		if num&1 == 1 {
			num = num ^ 1
			continue
		}

		num = num >> 1
	}

	return count
}

// 1111 1101
// 0000 0001

func main() {
	fmt.Println(numberOfSteps(13))
}
