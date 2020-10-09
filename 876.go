package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	a := head
	b := head

	for b.Next != nil && b.Next.Next != nil {
		a = a.Next
		b = b.Next.Next
	}

	if b.Next == nil {
		return a
	}

	return a.Next
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	list := &ListNode{Val: arr[0], Next: nil}

	tmp := list
	for i := 1; i < len(arr); i++ {
		node := &ListNode{Val: arr[i], Next: nil}
		tmp.Next = node
		tmp = tmp.Next
	}

	fmt.Println(middleNode(list))
}
