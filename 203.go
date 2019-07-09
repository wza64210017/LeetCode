package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	head = &ListNode{Next: head}
	q := head
	for q != nil && q.Next != nil {
		if q.Next.Val == val {
			q.Next = q.Next.Next
			continue
		}

		q = q.Next
	}

	return head.Next
}

//if head.Next != nil && head.Next.Val == val {
//	return head.Next.Next
//}

func main() {
	arr := []int{1, 2}
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

	list = removeElements(list, 1)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
