package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func partition(head *listNode, x int) *listNode {
	if head == nil || head.Next == nil {
		return head
	}

	less := &listNode{}
	more := &listNode{}

	p, q := less, more
	for head != nil {
		if head.Val < x {
			p.Next = head
			p = p.Next
		} else {
			q.Next = head
			q = q.Next
		}
		head = head.Next
	}

	q.Next = nil
	p.Next = more.Next
	return less.Next
}

func main() {
	//arr := []int{1, 4, 3, 2, 5, 2}
	arr := []int{2, 1}
	list := &listNode{
		Val:  arr[0],
		Next: nil,
	}

	tmp := list
	for i := 1; i < len(arr); i++ {
		tmp.Next = &listNode{
			Val:  arr[i],
			Next: nil,
		}

		tmp = tmp.Next
	}

	list = partition(list, 2)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
