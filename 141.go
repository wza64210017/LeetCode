package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	if head == head.Next {
		return true
	}

	//p:慢  q:快
	p, q := head, head
	for q.Next != nil && p.Next != nil && q.Next.Next != nil {
		q = q.Next.Next
		p = p.Next
		if p == q {
			return true
		}
	}

	return false
}

func main() {
	arr := []int{3, 2, 0, 4}
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

	tmp.Next = list.Next
	fmt.Println(hasCycle(list))
}
