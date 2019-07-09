package main

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	//1 -> 2 -> 3 -> 4 -> 5 -> 6
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	cur := slow
	next := cur.Next
	for next != nil {
		tmp := next.Next
		next.Next = cur
		cur = next
		next = tmp
	}

	slow.Next = nil
	for head != nil && cur.Next != nil {
		tmp := cur
		cur = cur.Next

		tmp.Next = head.Next
		head.Next = tmp
		head = tmp.Next
	}
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

	reorderList(list)
	for list != nil {
		fmt.Println(list.Val)
		time.Sleep(time.Second)
		list = list.Next
	}
}
