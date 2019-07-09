package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	odd := head
	even := head.Next
	tmp := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
		odd.Next = tmp
	}

	return head
}

func pick(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Next.Next == nil {
		return nil
	}

	tmp := head
	tmp.Next = pick(head.Next.Next)

	return tmp
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

	list = oddEvenList(list)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
