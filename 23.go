package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return merge(lists)
}

func merge(list []*ListNode) *ListNode {
	length := len(list)
	half := length / 2

	if length == 1 {
		return list[0]
	}

	if length == 2 {
		l1, l2 := list[0], list[1]
		if l1 == nil {
			return l2
		}

		if l2 == nil {
			return l1
		}

		var cur, mergeList *ListNode
		if l1.Val < l2.Val {
			mergeList, cur, l1 = l1, l1, l1.Next
		} else {
			mergeList, cur, l2 = l2, l2, l2.Next
		}

		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				cur.Next, l1 = l1, l1.Next
			} else {
				cur.Next, l2 = l2, l2.Next
			}
			cur = cur.Next
		}

		if l1 != nil {
			cur.Next = l1
		}

		if l2 != nil {
			cur.Next = l2
		}
		
		return mergeList
	}

	return mergeKLists([]*ListNode{mergeKLists(list[:half]), mergeKLists(list[half:])})
}

func main() {
	arr := [][]int{
		{1, 4, 5},
		{1, 3, 4},
		{2, 6},
	}

	lists := make([]*ListNode, len(arr))
	for k, first := range arr {
		lists[k] = &ListNode{
			Val:  first[0],
			Next: nil,
		}

		tmp := lists[k]
		for i := 1; i < len(first); i++ {
			tmp.Next = &ListNode{
				Val:  first[i],
				Next: nil,
			}

			tmp = tmp.Next
		}
	}

	newList := mergeKLists(lists)
	for newList != nil {
		fmt.Println(newList.Val)
		newList = newList.Next
	}
}
