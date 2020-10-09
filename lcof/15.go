package main

import "fmt"

func hammingWeight(num uint32) int {
	n := 0
	for num > 0 {
		if num & 1 == 1{
			n++
		}
		num = num >> 1
	}

	return n
}

func main() {
	fmt.Println(hammingWeight(128))
}
