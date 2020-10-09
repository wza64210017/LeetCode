package main

import "fmt"

func hammingDistance(x int, y int) int {
	num := 0

	x = x ^ y
	for x != 0 {
		if x & 1 == 1{
			num++
		}
		x = x >> 1

	}

	return num
}

func main() {
	x := 93
	y := 73

	fmt.Println(hammingDistance(x, y))
}