package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	newList := reverseList(head.Next)
//	tmp := head.Next
//	head.Next = nil
//	tmp.Next = head
//
//	return newList
//}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}

	return pre
}

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

	list = reverseList(list)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
