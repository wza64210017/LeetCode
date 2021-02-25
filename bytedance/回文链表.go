package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = reverse(slow.Next)
	print(slow)
	for {
		if slow.Val != head.Val {
			return false
		}

		if slow.Next == nil {
			break
		}

		slow = slow.Next
		head = head.Next
	}

	return true
}

// 反转链表
func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := reverse(head.Next)
	tmp := head.Next
	head.Next = nil
	tmp.Next = head

	return node
}

func main() {
	list := generate([]int{1, 2, 3, 5, 4, 3, 2, 1})
	fmt.Println(isPalindrome(list))
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
