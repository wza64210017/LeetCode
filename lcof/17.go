package main

import "fmt"

func printNumbers(n int) []int {
	num := 0
	for n > 0 {
		num = num*10 + 9
		n--
	}

	arr := make([]int, num)
	for i := 0; i < num; i++ {
		arr[i] = i+1
	}

	return arr
}

func main() {
	fmt.Println(printNumbers(0))
}
