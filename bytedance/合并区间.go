package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ret := make([][]int, 0, len(intervals))
	for i := 0; i < len(intervals)-1; i++ {
		L, R := intervals[i], intervals[i+1]

		if L[1] < R[0] {
			ret = append(ret, L)
			continue
		}

		if L[0] <= R[0] {
			intervals[i+1][0] = L[0]
		}

		if L[0] > R[0] {
			intervals[i+1][0] = R[0]
		}

		if L[1] <= R[1] {
			intervals[i+1][1] = R[1]
		}

		if L[1] > R[1] {
			intervals[i+1][1] = L[1]
		}
	}

	return append(ret, intervals[len(intervals)-1])
}

func main() {
	//fmt.Println(merge([][]int{
	//	{1, 3},
	//	{2, 6},
	//	{8, 10},
	//	{15, 18},
	//}))
	//
	//fmt.Println(merge([][]int{
	//	{1, 4},
	//	{4, 5},
	//}))
	//
	//fmt.Println(merge([][]int{}))
	//
	//fmt.Println(merge([][]int{
	//	{1, 4},
	//	{0, 4},
	//}))
	//
	//fmt.Println(merge([][]int{
	//	{1, 4},
	//	{0, 1},
	//}))
	//
	//fmt.Println(merge([][]int{
	//	{1, 4},
	//	{0, 0},
	//}))

	fmt.Println(merge([][]int{
		{2, 3},
		{4, 9},
		{6, 7},
		{8, 9},
		{1, 10},
	}))
}
