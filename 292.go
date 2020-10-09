package main

import "fmt"

//func canWinNim(n int) bool {
//	if n <= 0 {
//		return false
//	}
//
//	if n <= 3 {
//		return true
//	}
//
//	return !canWinNim(n - 1) || !canWinNim(n - 2) || !canWinNim(n - 3)
//}

func canWinNim(n int) bool {
	return n % 4 != 0
}

func main() {
	n := 4
	fmt.Println(canWinNim(n))

	//f(n)
	//1    n=1
	//2    n=2
	//3    n=3
	//f(n-1)+f(n-2)+f(n-3)     n>=4
}
