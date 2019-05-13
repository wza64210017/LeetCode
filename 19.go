package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	p, q := head, head

	i := 1
	for i <= n && q.Next != nil {
		q = q.Next
		i++
	}

	if q.Next == nil && i == n {
		head = head.Next
		return head
	}

	for q.Next != nil {
		q = q.Next
		p = p.Next
	}

	p.Next = p.Next.Next

	return head
}

func main() {
	arr1 := []int{1,2,3}
	list1 := &ListNode{Val: arr1[0], Next: nil}

	tmp := list1
	for i := 1; i < len(arr1); i++ {
		node := &ListNode{Val: arr1[i], Next: nil}
		tmp.Next = node
		tmp = tmp.Next
	}

	list := removeNthFromEnd(list1, 3)
	for list != nil && list.Next != nil {
		fmt.Println(list.Val)
		list = list.Next
	}

	if list != nil {
		fmt.Println(list.Val)
	}
}
