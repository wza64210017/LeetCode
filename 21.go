package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	newList := new(ListNode)
	if l1.Val < l2.Val {
		newList = &ListNode{Val: l1.Val, Next: nil}
		l1 = l1.Next
	} else {
		newList = &ListNode{Val: l2.Val, Next: nil}
		l2 = l2.Next
	}

	node := newList
	for l1 != nil && l2 != nil {
		var value int
		if l1.Val < l2.Val {
			value = l1.Val
			l1 = l1.Next
		} else {
			value = l2.Val
			l2 = l2.Next
		}

		node.Next = &ListNode{Val: value, Next: nil}
		node = node.Next
	}

	if l1 == nil {
		for l2 != nil {
			node.Next = &ListNode{Val: l2.Val, Next: nil}
			node = node.Next
			l2 = l2.Next
		}
	}

	if l2 == nil {
		for l1 != nil {
			node.Next = &ListNode{Val: l1.Val, Next: nil}
			node = node.Next
			l1 = l1.Next
		}
	}

	return newList
}

func main() {
	arr1 := []int{-9, 3}
	list1 := &ListNode{Val: arr1[0], Next: nil}

	tmp := list1
	for i := 1; i < len(arr1); i++ {
		node := &ListNode{Val: arr1[i], Next: nil}
		tmp.Next = node
		tmp = tmp.Next
	}

	arr2 := []int{5, 7}
	list2 := &ListNode{Val: arr2[0], Next: nil}

	tmp = list2
	for i := 1; i < len(arr2); i++ {
		node := &ListNode{Val: arr2[i], Next: nil}
		tmp.Next = node
		tmp = tmp.Next
	}

	list := mergeTwoLists(list1, list2)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
