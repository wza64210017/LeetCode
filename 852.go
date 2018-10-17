package main

import "fmt"

func peakIndexInMountainArray(A []int) int {
	for i := 1; i < len(A)-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			return i
		}
	}

	return 0
}

//山脉数组的峰顶索引
func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 2, 1, 0}))
}
