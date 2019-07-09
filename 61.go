package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func rotateRight(head *listNode, k int) *listNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	p := head
	length := 1
	for head.Next != nil {
		length++
		head = head.Next
	}

	head.Next = p
	move := length - k%length
	for move >= 0 {
		head = head.Next
		move--
	}

	for p.Next != head {
		p = p.Next
	}

	p.Next = nil
	return head
}

func main() {
	arr := []int{0, 1, 2}
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

	list = rotateRight(list, 4)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
