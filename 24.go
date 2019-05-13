package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func swapPairs(head *listNode) *listNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmp := head.Next
	head.Next = swapPairs(head.Next.Next)
	tmp.Next = head

	return tmp
}

func main() {
	arr := []int{1, 2, 3, 4}
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

	list = swapPairs(list)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}

}
