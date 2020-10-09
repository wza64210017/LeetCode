package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	for n != 1 {
		if n&1 != 0 {
			return false
		}

		n = n >> 1
	}

	return true
}

func main() {
	n := 0
	fmt.Println(isPowerOfTwo(n))
}
