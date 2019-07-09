package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

//func deleteDuplicates(head *listNode) *listNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	if head.Val == head.Next.Val {
//		for head.Next != nil && head.Val == head.Next.Val {
//			head = head.Next
//		}
//		return deleteDuplicates(head.Next)
//	}
//
//	head.Next = deleteDuplicates(head.Next)
//	return head
//}

func deleteDuplicates(head *listNode) *listNode {
	if head == nil || head.Next == nil {
		return head
	}

	x := &listNode{
		Val: 0,
		Next:head,
	}

	pre := x
	pos := head.Next
	flag := false

	for pos != nil {
		//若相邻值相等，pos指针后移，并将flag改成true
		if pos.Val == head.Val {
			flag = true
			pos = pos.Next
			continue
		}

		//flag 为true，修改pre的后继为pos
		if flag == true {
			pre.Next = pos
		} else {
			//flag为true，修改pre为head
			pre = head
		}

		flag = false
		head = pos
		pos = pos.Next
	}

	if head.Next != nil && head.Val == head.Next.Val {
		pre.Next = nil
	}

	return x.Next
}

func main() {
	arr := []int{1, 1, 3, 3, 3, 4, 5}
	list := &listNode{
		Val:  arr[0],
		Next: nil,
	}

	tmp := list
	for i := 1; i < len(arr); i++ {
		tmp.Next = &listNode{
			Val:  arr[i],
			Next: nil,
		}

		tmp = tmp.Next
	}

	list = deleteDuplicates(list)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
