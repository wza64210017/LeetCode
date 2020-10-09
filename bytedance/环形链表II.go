package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head
	isCircular := false
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			isCircular = true
			break
		}
	}

	fast = head
	for fast.Next != nil && slow.Next != nil && isCircular {
		if fast == slow {
			return fast
		}

		fast = fast.Next
		slow = slow.Next
	}

	return nil
}

func main() {
	arr := []int{1, 2}
	list := generate(arr)

	//node := list.Next
	//last := list
	//for last.Next != nil {
	//	last = last.Next
	//}
	//
	//last.Next = node
	fmt.Println(detectCycle(list))
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
