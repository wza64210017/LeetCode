package main

import (
	"fmt"
	"math"
)

func twoSum1(n int) []float64 {
	ret := make([]float64, 0, 5*n+1)    // 结果集
	num := math.Pow(6.0, float64(n)) // 总次数

	dp := make([]int, 6*n+1)
	for i := 1; i <= 6; i++ {
		dp[i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 6 * i; j >= i; j-- {
			dp[j] = 0
			for k := 1; k <= 6; k++ {
				if j-k >= i-1 {
					dp[j] += dp[j-k]
				}
			}
		}
	}

	for k := n; k <= 6*n; k++ {
		ret = append(ret, float64(dp[k])/num)
	}

	return ret
}

func main() {
	fmt.Println(twoSum1(2))
}
