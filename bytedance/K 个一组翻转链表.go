package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}

	i := 1
	p := head
	for i <= k && p != nil {
		p = p.Next
		i++
	}

	if i <= k {
		return head
	}

	newList := reverse(head, p)
	head.Next = reverseKGroup(p, k)

	return newList
}

func reverse(head, tail *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next == tail {
		return head
	}

	node := reverse(head.Next, tail)
	tmp := head.Next
	head.Next = nil
	tmp.Next = head

	return node
}

func main() {
	list := generate([]int{1, 2, 3, 4, 5})
	list = reverseKGroup(list, 2)

	print(list)
}

func generate(arr []int) *ListNode {
	list := &ListNode{
		Val:  arr[0],
		Next: nil,
	}

	p := list
	for i := 1; i < len(arr); i++ {
		p.Next = &ListNode{
			Val:  arr[i],
			Next: nil,
		}

		p = p.Next
	}

	return list
}

func print(list *ListNode) {
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
