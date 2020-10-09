package main

import "fmt"

func climbStairs(n int) int {
	m := make(map[int]int, n)
	m[1] = 1
	m[2] = 2

	for i := 3; i <= n; i++ {
		m[i] = m[i-1] + m[i-2]
		fmt.Println(i, m[i])
	}

	return m[n]
}

func main() {
	n := 4
	fmt.Println(climbStairs(n))
}