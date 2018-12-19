package main

import "fmt"

//type ListNode2 struct {
//	Val  int
//	Next *ListNode2
//}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	add := 0
	tmp := l1.Val + l2.Val
	if tmp >= 10 {
		tmp -= 10
		add = 1
	}

	l3 := &ListNode{Val: tmp, Next: nil}
	l1 = l1.Next
	l2 = l2.Next
	node := l3

	for l1 != nil && l2 != nil {
		tmp = l1.Val + l2.Val + add
		if tmp >= 10 {
			add = 1
			tmp -= 10
		} else {
			add = 0
		}

		node.Next = &ListNode{Val: tmp, Next: nil}
		node = node.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	if l2 != nil {
		for l2 != nil {
			tmp = l2.Val + add
			if tmp >= 10 {
				tmp -= 10
				add = 1
			} else {
				add = 0
			}

			node.Next = &ListNode{Val: tmp, Next: nil}
			node = node.Next
			l2 = l2.Next
		}
	}

	if l1 != nil {
		for l1 != nil {
			tmp = l1.Val + add
			if tmp >= 10 {
				tmp -= 10
				add = 1
			} else {
				add = 0
			}

			node.Next = &ListNode{Val: tmp, Next: nil}
			node = node.Next
			l1 = l1.Next
		}
	}

	if add == 1 {
		node.Next = &ListNode{Val: 1, Next: nil}
	}

	return l3
}

func main() {
	arr1 := []int{2, 4, 3}
	list1 := &ListNode{Val: arr1[0], Next: nil}

	tmp := list1
	for i := 1; i < len(arr1); i++ {
		node := &ListNode{Val: arr1[i], Next: nil}
		tmp.Next = node
		tmp = tmp.Next
	}

	arr2 := []int{5, 6, 7}
	list2 := &ListNode{Val: arr2[0], Next: nil}
	tmp = list2
	for i := 1; i < len(arr2); i++ {
		node := &ListNode{Val: arr2[i], Next: nil}
		tmp.Next = node
		tmp = tmp.Next
	}

	list := addTwoNumbers(list1, list2)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
