package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func deleteDuplicates(head *listNode) *listNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	for p.Next != nil {
		if p.Val != p.Next.Val {
			p = p.Next
			continue
		}

		//删除元素
		p.Next = p.Next.Next
	}

	return head
}

func main() {
	arr := []int{1, 1, 2, 3, 3}
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
