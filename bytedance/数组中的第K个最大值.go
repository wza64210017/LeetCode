package main

import (
	"container/heap"
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	h := H(nums)
	heap.Init(&h)

	num := -1
	for k > 0 {
		num = heap.Pop(&h).(int)
		k--
	}

	return num
}

type H []int

func (h H) Len() int {
	return len(h)
}

func (h H) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h H) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *H) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func (h *H) Pop() interface{} {
	tmp := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return tmp
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
