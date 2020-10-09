package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	node := l3
	add := 0

	for l1 != nil && l2 != nil {
		node.Next = &ListNode{}
		node = node.Next

		node.Val, add = calAdd(l1.Val + l2.Val + add)
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		node.Next = &ListNode{}
		node = node.Next

		node.Val, add = calAdd(l1.Val + node.Val + add)
		l1 = l1.Next
	}

	for l2 != nil {
		node.Next = &ListNode{}
		node = node.Next

		node.Val, add = calAdd(l2.Val + node.Val + add)
		l2 = l2.Next
	}

	for add > 0 {
		node.Next = &ListNode{}
		node = node.Next

		node.Val = add % 10
		add /= 10
	}

	return l3.Next
}

func calAdd(add int) (int, int) {
	return add % 10, add / 10
}

func main() {
	l1 := generate([]int{1, 7, 4})
	l2 := generate([]int{2, 3})

	l3 := addTwoNumbers(l1, l2)
	print(l3)
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
