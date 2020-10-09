package main

import "fmt"

const mod = 1000000007

//func fib(n int) int {
//	if n <= 1 {
//		return n
//	}
//
//	sum := fib(n-1) + fib(n-2)
//	if sum > mod {
//		return sum % mod
//	}
//
//	return sum
//}

func fib(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
		if a > mod {
			a = a % mod
		}
	}

	return a % mod
}

func main() {
	fmt.Println(fib(95))
}
