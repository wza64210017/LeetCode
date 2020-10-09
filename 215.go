package main

import (
	"container/heap"
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	p := P(nums)
	heap.Init(&p)

	num := -1
	for k > 0 {
		num = heap.Pop(&p).(int)
		k--
	}

	return num
}

type P []int

func (p P) Len() int {
	return len(p)
}

func (p P) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p P) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *P) Push(x interface{}) {
	*p = append(*p, x.(int))
}

func (p *P) Pop() interface{} {
	temp := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return temp
}

func main() {
	nums := []int{3,2,3,1,2,4,5,5,6}
	k := 4

	fmt.Println(findKthLargest(nums, k))
}
