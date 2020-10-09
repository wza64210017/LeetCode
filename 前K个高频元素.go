package main

import (
	"container/heap"
	"fmt"
)

type node struct {
	value int
	count int
}

type H []node

func (h *H) Len() int {
	return len(*h)
}

func (h H) Less(i, j int) bool  { return h[i].count < h[j].count }
func (h H) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *H) Push(i interface{}) { *h = append(*h, i.(node)) }

func (h *H) Pop() interface{} {
	tmp := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]

	return tmp
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num]++
	}

	h := &H{}
	heap.Init(h)

	for num, count := range m {
		heap.Push(h, node{value: num, count: count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ret := make([]int, k)
	for k > 0 {
		ret[k-1] = heap.Pop(h).(node).value
		k--
	}

	return ret
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 4}, 2))
}
