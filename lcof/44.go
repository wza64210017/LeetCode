package main

import (
	"fmt"
	"strconv"
)

const T = 10

func findNthDigit(n int) int {

	pos := 0
	count := 1
	for {
		// 位于当前数列中
		if n < count {
			//n = n - count/T - 1
			fmt.Println("n", n)

			a := strconv.Itoa(n / pos)
			b := n % pos

			fmt.Println(a, b)
			if b == 0 {
				return int(a[len(a)-1] - '0')
			}

			return int(a[b-1] - '0')
		}

		pos++
		count *= T
		fmt.Println(count/T)
		n = n - count/T - 1
	}
}

func main() {
	fmt.Println(findNthDigit(11))
}
