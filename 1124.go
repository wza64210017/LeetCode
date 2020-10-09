package main

import "fmt"

func longestWPI(hours []int) int {

	max := 0
	for i := len(hours) - 1; i >= 0; i-- {
		flag := 0
		for j := i; j >= 0; j-- {
			if hours[j] > 8 {
				flag++
			} else {
				flag--
			}

			if flag > 0 && i-j+1 > max {
				max = i - j + 1
			}
		}
	}

	return max
}

func main() {
	hours := []int{6, 9, 9, 9}
	fmt.Println(longestWPI(hours))
}
