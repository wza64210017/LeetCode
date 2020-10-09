package main

import "fmt"

func add(a int, b int) int {
	for b != 0 {
		fmt.Println("a^b", a^b)
		fmt.Println("a&b", a&b, "<<1", (a&b)<<1)

		tmp := a ^ b
		b = (a & b) << 1
		a = tmp
	}

	return a
}

func main() {
	// 0000 0011
	// 0000 1001
	fmt.Println(add(3, 9))
}
