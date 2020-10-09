package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type H []int

func (h H) Len() int           { return len(h) }
func (h H) Less(i, j int) bool { return h[i] < h[j] }
func (h H) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *H) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func (h *H) Pop() interface{} {
	t := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return t
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := make(H, 0)
	heap.Init(&h)

	for _, node := range lists {
		for node != nil {
			heap.Push(&h, node.Val)
			node = node.Next
		}
	}

	node := &ListNode{}
	p := node
	for h.Len() > 0 {
		p.Next = &ListNode{
			Val: heap.Pop(&h).(int),
		}

		p = p.Next
	}

	return node.Next
}
