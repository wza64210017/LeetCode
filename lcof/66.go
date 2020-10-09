package main

import "fmt"

func constructArr(a []int) []int {
	length := len(a)
	if length == 0 {
		return nil
	}

	b := make([]int, length)
	b[0] = 1
	tmp := 1

	for i := 1; i < length; i++ {
		b[i] = b[i-1] * a[i-1]
	}

	for i := length - 2; i >= 0; i-- {
		tmp *= a[i+1]
		b[i] *= tmp
	}

	return b
}

func main() {
	fmt.Println(constructArr([]int{1, 2, 3, 4, 5}))
}
