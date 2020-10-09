package main

import "fmt"

// 单链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(arr []int) *ListNode {
	head := &ListNode{
		Val: arr[0],
	}

	p := head
	for i := 1; i < len(arr); i++ {
		p.Next = &ListNode{
			Val: arr[i],
		}

		p = p.Next
	}

	return head
}

// 打印链表
func printList(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}