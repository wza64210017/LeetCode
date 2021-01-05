package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	var ret float64 = 1
	for n > 0 {
		if n%2 == 1 {
			ret *= x
		}

		x = x * x
		n >>= 1
	}

	return ret
}

func main() {
	fmt.Println(myPow(2.10000, 3))
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPow(2.00000, -2))
	fmt.Println(myPow(0.44528, 0))
	fmt.Println(myPow(0.00001, 2147483647))
}
