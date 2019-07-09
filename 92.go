package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	p := dummy
	for i := 1; i < m; i++ {
		p = p.Next
	}

	head = p.Next
	for i := m; i < n; i++ {
		next := head.Next
		head.Next = next.Next
		next.Next = head
		p.Next = next
	}

	return dummy.Next
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	list := &ListNode{
		Val:  arr[0],
		Next: nil,
	}

	tmp := list
	for i := 1; i < len(arr); i++ {
		tmp.Next = &ListNode{
			Val:  arr[i],
			Next: nil,
		}

		tmp = tmp.Next
	}

	list = reverseBetween(list, 2, 4)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
