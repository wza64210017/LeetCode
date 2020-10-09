package main

import (
	"container/heap"
	"fmt"
)

type node struct {
	s   string
	num int
}

type H []node

func (h H) Len() int {
	return len(h)
}

func (h H) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h H) Less(i int, j int) bool {
	if h[i].num == h[j].num {
		return h[i].s > h[j].s
	}

	return h[i].num < h[j].num
}

func (h *H) Push(i interface{}) {
	*h = append(*h, i.(node))
}

func (h *H) Pop() interface{} {
	tmp := (*h)[0]

	old := *h
	old[0] = (old)[h.Len()-1]
	old = old[:h.Len()-1]
	*h = old

	return tmp
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int, len(words))
	for _, word := range words {
		m[word]++
	}

	p := make(H, 0, len(m))
	heap.Init(&p)

	for s, num := range m {
		p.Push(node{s: s, num: num})
		if p.Len() > k {
			p.Pop()
		}
	}

	i := k-1
	ret := make([]string, k)
	for i > 0 {
		ret[i] = p.Pop().(node).s
		i--
	}

	return ret
}

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
}
