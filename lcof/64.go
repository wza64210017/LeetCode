package main

import "fmt"

var ret int

func sumNums(n int) int {
	ret = 0
	cal(n)
	return ret
}

func cal(n int) bool {
	ret += n
	return n > 0 && cal(n-1)
}

func main() {
	fmt.Println(sumNums(3))
	fmt.Println(sumNums(2))
	fmt.Println(sumNums(4))
}
