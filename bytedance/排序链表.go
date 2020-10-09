package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	var h H
	heap.Init(&h)

	p := head
	for p != nil {
		heap.Push(&h, p.Val)
		p = p.Next
	}

	n := head
	for h.Len() > 0 {
		n.Val = heap.Pop(&h).(int)
	}

	return head
}

type H []int

func (h H) Len() int           { return len(h) }
func (h H) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h H) Less(i, j int) bool { return h[i] < h[j] }

func (h *H) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func (h *H) Pop() interface{} {
	tmp := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return tmp
}
