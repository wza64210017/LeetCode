package main

import "fmt"

const mod = 1000000007

func numWays(n int) int {
	m := make(map[int]int, n)
	m[0], m[1] = 1, 1

	for i := 2; i <= n; i++ {
		m[i] = (m[i-1] + m[i-2]) % mod
	}

	return m[n]
}

func main() {
	fmt.Println(numWays(7))
}
