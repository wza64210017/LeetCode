package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//func detectCycle(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return nil
//	}
//
//	m := make(map[*ListNode]bool)
//	for head != nil {
//		if m[head] == true {
//			return head
//		}
//
//		m[head] = true
//		head = head.Next
//	}
//
//	return nil
//}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	cycle := false
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		fmt.Println("slow", slow)
		fmt.Println("fast", fast)
		if slow  == fast {
			cycle = true
			break
		}
	}

	return nil

	if cycle {
		cur := head
		for cur != slow {
			cur = cur.Next
			slow = slow.Next
		}
		return cur
	}
	return nil
}

func main() {
	arr := []int{3,2,0,-4, 5}
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

	tmp.Next = list.Next.Next
	fmt.Println(detectCycle(list))
}
