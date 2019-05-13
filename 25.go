package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func reverseKGroup(head *listNode, k int) *listNode {
	if k <= 1 {
		return head
	}

	node := head
	i := 1
	for i <= k && node != nil {
		node = node.Next
		i++
	}

	if i <= k {
		return head
	}

	//1 -> 2 -> 3 -> 4 -> 5
	//k = 3
	//此时head = 1
	//1 -> 2 -> 3  -->  4 -> 5  利用递归拆为
	//1 -> 2 -> 3 翻转 3 -> 2 -> 1
	//1(head)变成尾节点
	//head续接下一段(递归)

	newList := reverse(head, node)
	head.Next = reverseKGroup(node, k)

	return newList
}

func reverse(head *listNode, tail *listNode) *listNode {
	if head == nil || head.Next == nil || head.Next == tail {
		return head
	}

	newList := reverse(head.Next, tail)
	tmp := head.Next
	tmp.Next = head
	head.Next = nil

	return newList
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
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

	list = reverseKGroup(list, 2)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}

}
